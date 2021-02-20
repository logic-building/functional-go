package basic

import "strings"

// FilterMapIONumberPtrTest is template to generate itself for different combination of data type.
func FilterMapIONumberPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : some logic
	var v1 <INPUT_TYPE> = 1
	var v2 <INPUT_TYPE> = 2
	var v3 <INPUT_TYPE> = 3

	var vo3 <OUTPUT_TYPE> = 3
	var vo4 <OUTPUT_TYPE> = 4

	expectedList := []*<OUTPUT_TYPE>{&vo3, &vo4}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&v1, &v2, &v3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num != 1
}
func plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) *<OUTPUT_TYPE> {
	c := <OUTPUT_TYPE>(*num + 1)
	return &c
}
`
}

// FilterMapIOStrNumberPtrTest is template to generate itself for different combination of data type.
func FilterMapIOStrNumberPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = 10

	var vOne <INPUT_TYPE> = "one"
	var vTen <INPUT_TYPE> = "ten"

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrPtr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrPtr, []*<INPUT_TYPE>{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrPtr(num *<INPUT_TYPE>) bool {
	return *num != "one"
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrPtr(num *string) *<OUTPUT_TYPE> {
	var r <OUTPUT_TYPE> = <OUTPUT_TYPE>(0)
	if *num == "ten" {
		r = <OUTPUT_TYPE>(10)
		return &r
	}
	return &r
}
`
}

// FilterMapIONumberStrPtrTest is template to generate itself for different combination of data type.
func FilterMapIONumberStrPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var ov10 <OUTPUT_TYPE> = "10"
	var iv1 <INPUT_TYPE> = 1
	var iv10 <INPUT_TYPE> = 10
	expectedList := []*<OUTPUT_TYPE>{&ov10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumPtr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumPtr, []*<INPUT_TYPE>{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumPtr(num *<INPUT_TYPE>) bool {
	return *num != 1
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumPtr(num *<INPUT_TYPE>) *<OUTPUT_TYPE> {
	var r <OUTPUT_TYPE> = <OUTPUT_TYPE>(0)
	if *num == 10 {
		r = "10"
		return &r
	}
	return &r
}
`
}

// FilterMapIONumberBoolPtrTest is template to generate itself for different combination of data type.
func FilterMapIONumberBoolPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vto <OUTPUT_TYPE> = true
	var vfo <OUTPUT_TYPE> = false

	var vi1 <INPUT_TYPE> = 1
	var vi10 <INPUT_TYPE> = 10
	var vi0 <INPUT_TYPE>

	expectedList := []*<OUTPUT_TYPE>{&vto, &vfo}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num != 1
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) *bool {
	r := *num > 0
	return &r
}
`
}

// FilterMapIOStrBoolPtrTest is template to generate itself for different combination of data type.
func FilterMapIOStrBoolPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic

	var vto <OUTPUT_TYPE> = true
	var vfo <OUTPUT_TYPE> = false

	var vi1 <INPUT_TYPE> = "1"
	var vi10 <INPUT_TYPE> = "10"
	var vi0 <INPUT_TYPE> = "0"

	expectedList := []*<OUTPUT_TYPE>{&vto, &vfo}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num != "1"
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) *bool {
	var t bool = true
	var f bool = false

	if *num == "10" {
		return &t
	}
	return &f
}
`
}

// FilterMapIOBoolNumberPtrTest is template to generate itself for different combination of data type.
func FilterMapIOBoolNumberPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic

	var vo10 <OUTPUT_TYPE> = 10

	var vit <INPUT_TYPE> = true
	var vif <INPUT_TYPE> = false

	expectedList := []*<OUTPUT_TYPE>{&vo10, &vo10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num == true
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *bool) *<OUTPUT_TYPE> {
	var v10 <OUTPUT_TYPE> = 10
	var v0 <OUTPUT_TYPE> = 0

	if *num == true {
		return &v10
	}
	return &v0
}
`
}

// FilterMapIOBoolStrPtrTest is template to generate itself for different combination of data type.
func FilterMapIOBoolStrPtrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = "10"

	var vit <INPUT_TYPE> = true
	var vif <INPUT_TYPE> = false

	expectedList := []*<OUTPUT_TYPE>{&vo10, &vo10}
	newList := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vit, &vit, &vif})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) bool {
	return *num == true
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *bool) *<OUTPUT_TYPE> {
	var v10 <OUTPUT_TYPE> = "10"
	var v0 <OUTPUT_TYPE> = "0"

	if *num == true {
		return &v10
	}
	return &v0
}
`
}

//**********************************Err***********************************

// FilterMapIONumberPtrErrTest is template to generate itself for different combination of data type.
func FilterMapIONumberPtrErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : some logic
	var v1 <INPUT_TYPE> = 1
	var v2 <INPUT_TYPE> = 2
	var v3 <INPUT_TYPE> = 3
	var v4 <INPUT_TYPE> = 4
	var v5 <INPUT_TYPE> = 5

	var vo5 <OUTPUT_TYPE> = 5
	var vo6 <OUTPUT_TYPE> = 6

	expectedList := []*<OUTPUT_TYPE>{&vo5, &vo6}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, plusOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v1, &v4, &v5})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, []*<INPUT_TYPE>{})

	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, plusOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v2, &v3})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, plusOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v3})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (bool, error) {
	if *num == 2 {
		return false, errors.New(" 2 is not valid number for this test")
	}
	return *num != 1, nil
}
func plusOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (*<OUTPUT_TYPE>, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	c := <OUTPUT_TYPE>(*num + 1)
	return &c, nil
}
`
}

// FilterMapIOStrNumberPtrErrTest is template to generate itself for different combination of data type.
func FilterMapIOStrNumberPtrErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = 10

	var vOne <INPUT_TYPE> = "one"
	var vTwo <INPUT_TYPE> = "two"
	var vThree <INPUT_TYPE> = "three"
	var vTen <INPUT_TYPE> = "ten"

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrPtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrPtrErr, []*<INPUT_TYPE>{&vOne, &vTen})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, nil)

	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrPtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrPtrErr, []*<INPUT_TYPE>{&vTwo, &vThree})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrPtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrPtrErr, []*<INPUT_TYPE>{&vThree})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>StrPtrErr(num *<INPUT_TYPE>) (bool, error) {
	if *num == "two" {
		return false, errors.New("Two is not valid for this test")
	}
	return *num != "one", nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>StrPtrErr(num *string) (*<OUTPUT_TYPE>, error) {
	var r <OUTPUT_TYPE> = <OUTPUT_TYPE>(0)
	if *num == "three" {
		return nil, errors.New("three is not valid value for this test")
	}
	if *num == "ten" {
		r = <OUTPUT_TYPE>(10)
		return &r, nil
	}
	return &r, nil
}
`
}

// FilterMapIONumberStrPtrErrTest is template to generate itself for different combination of data type.
func FilterMapIONumberStrPtrErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var ov10 <OUTPUT_TYPE> = "10"
	var iv1 <INPUT_TYPE> = 1
	var iv2 <INPUT_TYPE> = 2
	var iv3 <INPUT_TYPE> = 3
	var iv10 <INPUT_TYPE> = 10
	expectedList := []*<OUTPUT_TYPE>{&ov10}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumPtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumPtrErr, []*<INPUT_TYPE>{&iv1, &iv10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumPtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumPtrErr, []*<INPUT_TYPE>{&iv1, &iv2, &iv10})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumPtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumPtrErr, []*<INPUT_TYPE>{&iv1, &iv3, &iv10})

	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
}

func notOne<FINPUT_TYPE><FOUTPUT_TYPE>NumPtrErr(num *<INPUT_TYPE>) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test")
	}
	return *num != 1, nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>NumPtrErr(num *<INPUT_TYPE>) (*<OUTPUT_TYPE>, error) {
	var r <OUTPUT_TYPE> = <OUTPUT_TYPE>(0)
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	if *num == 10 {
		r = "10"
		return &r, nil
	}
	return &r, nil
}
`
}

// FilterMapIONumberBoolPtrErrTest is template to generate itself for different combination of data type.
func FilterMapIONumberBoolPtrErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vto <OUTPUT_TYPE> = true
	var vfo <OUTPUT_TYPE> = false

	var vi1 <INPUT_TYPE> = 1
	var vi2 <INPUT_TYPE> = 2
	var vi3 <INPUT_TYPE> = 3
	var vi10 <INPUT_TYPE> = 10
	var vi0 <INPUT_TYPE> = 0

	expectedList := []*<OUTPUT_TYPE>{&vto, &vfo}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi2, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi3, &vi10, &vi0})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (bool, error) {
	if *num == 2 {
		return false, errors.New("2 is not valid number for this test") 
	}
	return *num != 1, nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (*bool, error) {
	if *num == 3 {
		return nil, errors.New("3 is not valid number for this test")
	}
	r := *num > 0
	return &r, nil
}
`
}

// FilterMapIOStrBoolPtrErrTest is template to generate itself for different combination of data type.
func FilterMapIOStrBoolPtrErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic

	var vto <OUTPUT_TYPE> = true
	var vfo <OUTPUT_TYPE> = false

	var vi1 <INPUT_TYPE> = "1"
	var vi2 <INPUT_TYPE> = "2"
	var vi3 <INPUT_TYPE> = "3"
	var vi10 <INPUT_TYPE> = "10"
	var vi0 <INPUT_TYPE> = "0"

	expectedList := []*<OUTPUT_TYPE>{&vto, &vfo}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi1, &vi10, &vi0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi1, &vi10, &vi2})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi1, &vi10, &vi3})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (bool, error) {
	if *num == "2" {
		return false, errors.New("2 is not valid value for this test")
	}
	return *num != "1", nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (*bool, error) {
	if *num == "3" {
		return nil, errors.New("3 is not valid value for this test")
	}
	var t bool = true
	var f bool = false

	if *num == "10" {
		return &t, nil
	} 
	return &f, nil
}
`
}

// FilterMapIOBoolNumberPtrErrTest is template to generate itself for different combination of data type.
func FilterMapIOBoolNumberPtrErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic

	var vo10 <OUTPUT_TYPE> = 10
	
	var vit <INPUT_TYPE> = true
	var vif <INPUT_TYPE> = false

	expectedList := []*<OUTPUT_TYPE>{&vo10, &vo10}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr2, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr2(num *<INPUT_TYPE>) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *bool) (*<OUTPUT_TYPE>, error) {
	if *num == false {
		return nil, errors.New("false is error for this test")
	}
	var v10 <OUTPUT_TYPE> = 10
	var v0 <OUTPUT_TYPE> = 0

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}
`
}

// FilterMapIOBoolStrPtrErrTest is template to generate itself for different combination of data type.
func FilterMapIOBoolStrPtrErrTest() string {
	return `
func TestFilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = "10"
	
	var vit <INPUT_TYPE> = true
	var vif <INPUT_TYPE> = false

	expectedList := []*<OUTPUT_TYPE>{&vo10, &vo10}
	newList, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vit, &vit})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vit, &vit, nil})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err = FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr2, someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vit, &vit, &vif})
	if err == nil {
		t.Errorf("FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
}
func notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return *num == true, nil
}

func notOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr2(num *<INPUT_TYPE>) (bool, error) {
	if num == nil {
		return false, errors.New("nil is error in this test")
	}
	return true, nil
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *bool) (*<OUTPUT_TYPE>, error) {
	if *num == false {
		return nil, errors.New("false is error in this test")
	}
	var v10 <OUTPUT_TYPE> = "10"
	var v0 <OUTPUT_TYPE> = "0"

	if *num == true {
		return &v10, nil
	}
	return &v0, nil
}
`
}

// ReplaceActivityFilterMapIOPtrErr replaces ...
func ReplaceActivityFilterMapIOPtrErr(code string) string {
	s1 := `_ "errors"
	"reflect"
	"testing"
)

func TestFilterMapIntInt64PtrErr(t *testing.T) {`
	s2 := `"errors"
	"testing"
)

func TestFilterMapIntInt64PtrErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)
	return code
}
