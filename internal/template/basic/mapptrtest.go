package basic

import "strings"

// MapPtrTest applies the function(1st argument) on each item of the list and returns new list
func MapPtrTest() string {
	return `
func TestMap<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v5 <TYPE> = 5
	var v6 <TYPE> = 6
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*<TYPE>{&v6, &v7, &v8}
	partialAdd<FTYPE>Ptr := func(num *<TYPE>) *<TYPE> {
		return add<FTYPE>Ptr(&v5, num)
	}
	sumList := Map<FTYPE>Ptr(partialAdd<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("Map<FTYPE>Ptr failed")
	}

	if len(Map<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func add<FTYPE>Ptr(num1, num2 *<TYPE>) *<TYPE> {
    result := *num1 + *num2
	return &result
}

`
}

// MapPtrBoolTest applies the function(1st argument) on each item of the list and returns new list
func MapPtrBoolTest() string {
	return `
func TestMap<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []*<TYPE>{&vf}
	
	newList := Map<FTYPE>Ptr(inverse<FTYPE>Ptr, []*<TYPE>{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("Map<FTYPE>Ptr failed")
	}

	if len(Map<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FTYPE>Ptr failed.")
	}
}

func inverse<FTYPE>Ptr(num1 *<TYPE>) *<TYPE> {
	vt := true
    if *num1 == true {
		v := false
		return &v
	} 
	return &vt
}

`
}

// MapPtrErrTest applies the function(1st argument) on each item of the list and returns new list
func MapPtrErrTest() string {
	return `
func TestMap<FTYPE>PtrErr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v5 <TYPE> = 5
	var v6 <TYPE> = 6
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8


	// Test: add 5 to each item in the list
	expectedSumList := []*<TYPE>{&v6, &v7, &v8}
	partialAdd<FTYPE>Ptr := func(num *<TYPE>) (*<TYPE>, error) {
		r, err := add<FTYPE>PtrErr(&v5, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	sumList, _ := Map<FTYPE>PtrErr(partialAdd<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("Map<FTYPE>PtrErr failed")
	}

	r, _ := Map<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FTYPE>PtrErr failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 <TYPE> = 0
	partialAdd<FTYPE>Ptr2 := func(num *<TYPE>) (*<TYPE>, error) {
		r, err := add<FTYPE>PtrErr(&v0, num)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	_, err := Map<FTYPE>PtrErr(partialAdd<FTYPE>Ptr2, []*<TYPE>{&v0})
	if err == nil {
		t.Errorf("Map<FTYPE>PtrErr failed")
	}

	
}

func add<FTYPE>PtrErr(num1, num2 *<TYPE>) (*<TYPE>, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}

`
}

// MapPtrErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func MapPtrErrBoolTest() string {
	return `
func TestMap<FTYPE>PtrErr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []*<TYPE>{&vf}
	
	newList, _ := Map<FTYPE>PtrErr(inverse<FTYPE>PtrErr, []*<TYPE>{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("Map<FTYPE>PtrErr failed")
	}

	_, err1 := Map<FTYPE>PtrErr(inverse<FTYPE>PtrErr, []*<TYPE>{&vf})
	if err1 == nil  {
		t.Errorf("Map<FTYPE>PtrErr failed")
	}

	r, _ := Map<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FTYPE>PtrErr failed.")
	}
}

func inverse<FTYPE>PtrErr(num1 *<TYPE>) (*<TYPE>, error) {
	if *num1 == false {
		return nil, errors.New("False is not allowed")
	}
	vt := true
    if *num1 == true {
		v := false
		return &v, nil
	} 
	return &vt, nil
}

`
}

// ReplaceActivityMapPtrErrTest replaces ...
func ReplaceActivityMapPtrErrTest(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMapIntPtrErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"reflect"
	"testing"
)

func TestMapIntPtrErr(t *testing.T) {`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func addStrPtrErr(num1, num2 *string) (*string, error) {
	if *num1 < 1 {
		return nil, errors.New("Negative value not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}`
	s2 = `func addStrPtrErr(num1, num2 *string) (*string, error) {
	if *num1 == "0" {
		return nil, errors.New("0 is not allowed")
	}
    result := *num1 + *num2
	return &result, nil
}`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestMapStrPtrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v5 string = "5"
	var v6 string = "6"
	var v7 string = "7"
	var v8 string = "8"


	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}`

	s2 = `func TestMapStrPtrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v5 string = "5"
	var v6 string = "51"
	var v7 string = "52"
	var v8 string = "53"


	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}

// MapErrTest applies the function(1st argument) on each item of the list and returns new list
func MapErrTest() string {
	return `
func TestMap<FTYPE>Err(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v5 <TYPE> = 5
	var v6 <TYPE> = 6
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8

	// Test: add 5 to each item in the list
	expectedSumList := []<TYPE>{v6, v7, v8}
	partialAdd<FTYPE> := func(num <TYPE>) (<TYPE>, error) {
		r, err := add<FTYPE>Err(v5, num)
		if err != nil {
			return <TYPE>(0), err
		}
		return r, nil
	}
	sumList, _ := Map<FTYPE>Err(partialAdd<FTYPE>, []<TYPE>{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("Map<FTYPE>Err failed")
	}

	r, _ := Map<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FTYPE>Err failed.")
		t.Errorf(reflect.String.String())
	}

	var v0 <TYPE> = 0
	partialAdd<FTYPE>2 := func(num <TYPE>) (<TYPE>, error) {
		r, err := add<FTYPE>Err(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}
	_, err := Map<FTYPE>Err(partialAdd<FTYPE>2, []<TYPE>{v0})
	if err == nil {
		t.Errorf("Map<FTYPE>Err failed")
	}
}

func add<FTYPE>Err(num1, num2 <TYPE>) (<TYPE>, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}
`
}

// MapErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func MapErrBoolTest() string {
	return `
func TestMap<FTYPE>Err(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []<TYPE>{vf}

	newList, _ := Map<FTYPE>Err(inverse<FTYPE>Err, []<TYPE>{vt})
	if newList[0] != expectedSumList[0] {
		t.Errorf("Map<FTYPE>Err failed")
	}

	_, err1 := Map<FTYPE>Err(inverse<FTYPE>Err, []<TYPE>{vf})
	if err1 == nil {
		t.Errorf("Map<FTYPE>Err failed")
	}

	r, _ := Map<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FTYPE>Err failed.")
	}
}

func inverse<FTYPE>Err(num1 <TYPE>) (<TYPE>, error) {
	if num1 == false {
		return false, errors.New("False is not allowed")
	}
	vt := true
	if num1 == true {
		v := false
		return v, nil
	}
	return vt, nil
}
`
}

// ReplaceActivityMapErrTest replaces ...
func ReplaceActivityMapErrTest(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMapIntErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"reflect"
	"testing"
)

func TestMapIntErr(t *testing.T) {`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func addStrErr(num1, num2 string) (string, error) {
	if num1 < 1 {
		return 0, errors.New("Negative value not allowed")
	}
	result := num1 + num2
	return result, nil
}`
	s2 = `func addStrErr(num1, num2 string) (string, error) {
	if num1 == "0" {
		return "", errors.New("0 is not allowed")
	}
	result := num1 + num2
	return result, nil
}`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestMapStrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v5 string = "5"
	var v6 string = "6"
	var v7 string = "7"
	var v8 string = "8"

	// Test: add 5 to each item in the list
	expectedSumList := []string{v6, v7, v8}`

	s2 = `func TestMapStrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v5 string = "5"
	var v6 string = "51"
	var v7 string = "52"
	var v8 string = "53"

	// Test: add 5 to each item in the list
	expectedSumList := []string{v6, v7, v8}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `var v0 string = "0"
	partialAddStr2 := func(num string) (string, error) {
		r, err := addStrErr(v0, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}`

	s2 = `var v0 string = "0"
	partialAddStr2 := func(num string) (string, error) {
		r, err := addStrErr(v0, num)
		if err != nil {
			return "0", err
		}
		return r, nil
	}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}
