package basic

import "strings"

// MapIONumberPtrTest is template to generate itself for different combination of data type.
func MapIONumberPtrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 <OUTPUT_TYPE> = 2
	var vo3 <OUTPUT_TYPE> = 3
	var vo4 <OUTPUT_TYPE> = 4

	var vi1 <INPUT_TYPE> = 1
	var vi2 <INPUT_TYPE> = 2
	var vi3 <INPUT_TYPE> = 3

	expectedList := []*<OUTPUT_TYPE>{&vo2, &vo3, &vo4}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrNumberPtrTest is template to generate itself for different combination of data type.
func MapIOStrNumberPtrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = 10

	var vi10 <INPUT_TYPE> = "10"

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) *<OUTPUT_TYPE> {
	var r <OUTPUT_TYPE>
	if *num == "10" {
		r = 10
	}
	return &r
}
`
}

// MapIONumberStrPtrTest is template to generate itself for different combination of data type.
func MapIONumberStrPtrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = "10"
	var vi10 <INPUT_TYPE> = 10

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(num *<INPUT_TYPE>) *string {
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r
}
`
}

// MapIONumberBoolPtrTest is template to generate itself for different combination of data type.
func MapIONumberBoolPtrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vt <OUTPUT_TYPE> = true
	var vf <OUTPUT_TYPE> = false

	var v10 <INPUT_TYPE> = 10
	var v0 <INPUT_TYPE>

	expectedList := []*<OUTPUT_TYPE>{&vt, &vf}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrBoolPtrTestPtr is template to generate itself for different combination of data type.
func MapIOStrBoolPtrTestPtr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	// Test : someLogic
	var vt <OUTPUT_TYPE> = true
	var vf <OUTPUT_TYPE> = false

	var v10 <INPUT_TYPE> = "10"
	var v0 <INPUT_TYPE> = "0"

	expectedList := []*<OUTPUT_TYPE>{&vt, &vf}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolNumberPtrTest is template to generate itself for different combination of data type.
func MapIOBoolNumberPtrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	var vt <INPUT_TYPE> = true
	var vf <INPUT_TYPE> = false

	var v10 <OUTPUT_TYPE> = 10
	var v0 <OUTPUT_TYPE>

	// Test : someLogic
	expectedList := []*<OUTPUT_TYPE>{&v10, &v0}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolStrPtrTest is template to generate itself for different combination of data type.
func MapIOBoolStrPtrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(t *testing.T) {
	var vt <INPUT_TYPE> = true
	var vf <INPUT_TYPE> = false

	var v10 <OUTPUT_TYPE> = "10"
	var v0 <OUTPUT_TYPE> = "0"

	// Test : someLogic
	expectedList := []*<OUTPUT_TYPE>{&v10, &v0}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Ptr, []*<INPUT_TYPE>{&vt, &vf})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(nil, []*<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIONumberPtrErrTest is template to generate itself for different combination of data type.
func MapIONumberPtrErrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 <OUTPUT_TYPE> = 2
	var vo3 <OUTPUT_TYPE> = 3

	var vi1 <INPUT_TYPE> = 1
	var vi2 <INPUT_TYPE> = 2
	var vi3 <INPUT_TYPE> = 3

	expectedList := []*<OUTPUT_TYPE>{&vo2, &vo3}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrNumberPtrErrTest is template to generate itself for different combination of data type.
func MapIOStrNumberPtrErrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = 10

	var vi10 <INPUT_TYPE> = "10"
	var vi0 <INPUT_TYPE> = "0"

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi10, &vi0})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (*<OUTPUT_TYPE>, error) {
	if *num == "0" {
		return nil, errors.New("0 is invalid for this test")
	}
	var r <OUTPUT_TYPE> = 0
	if *num == "10" {
		r = 10
	}
	return &r, nil
}
`
}

// MapIONumberStrPtrErrTest is template to generate itself for different combination of data type.
func MapIONumberStrPtrErrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 <OUTPUT_TYPE> = "10"
	var vi10 <INPUT_TYPE> = 10
	var vi0 <INPUT_TYPE>

	expectedList := []*<OUTPUT_TYPE>{&vo10}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vi10, &vi0})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(num *<INPUT_TYPE>) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0 is invalid")
	}
	r := "0"
	if *num == 10 {
		r = string("10")
	}
	return &r, nil
}
`
}

// MapIONumberBoolPtrErrTest is template to generate itself for different combination of data type.
func MapIONumberBoolPtrErrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vt <OUTPUT_TYPE> = true
	var vf <OUTPUT_TYPE> = false

	var v10 <INPUT_TYPE> = 10
	var v0 <INPUT_TYPE>
	var v3 <INPUT_TYPE> = 3
	
	expectedList := []*<OUTPUT_TYPE>{&vt, &vf}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v10, &v3})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrBoolPtrErrTest is template to generate itself for different combination of data type.
func MapIOStrBoolPtrErrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	// Test : someLogic
	var vt <OUTPUT_TYPE> = true
	var vf <OUTPUT_TYPE> = false

	var v10 <INPUT_TYPE> = "10"
	var v0 <INPUT_TYPE> = "0"
	var v3 <INPUT_TYPE> = "3"

	expectedList := []*<OUTPUT_TYPE>{&vt, &vf}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v10, &v0})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&v10, &v0, &v3})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolNumberPtrErrTest is template to generate itself for different combination of data type.
func MapIOBoolNumberPtrErrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	var vt <INPUT_TYPE> = true
	var vf <INPUT_TYPE> = false

	var v10 <OUTPUT_TYPE> = 10
	var v0 <OUTPUT_TYPE>

	// Test : someLogic
	expectedList := []*<OUTPUT_TYPE>{&v10, &v0}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vt, &vf})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolStrPtrErrTest is template to generate itself for different combination of data type.
func MapIOBoolStrPtrErrTest() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(t *testing.T) {
	var vt <INPUT_TYPE> = true
	var vf <INPUT_TYPE> = false

	var v10 <OUTPUT_TYPE> = "10"
	var v0 <OUTPUT_TYPE> = "0"

	// Test : someLogic
	expectedList := []*<OUTPUT_TYPE>{&v10, &v0}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(nil, []*<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr, []*<INPUT_TYPE>{&vt, &vf})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// ReplaceActivityMapIOPtrErr replaces ...
func ReplaceActivityMapIOPtrErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMapIntInt64PtrErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"reflect"
	"testing"
)

func TestMapIntInt64PtrErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)
	return code
}
