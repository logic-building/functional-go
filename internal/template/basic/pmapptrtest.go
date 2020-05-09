package basic

import "strings"

// PMap<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func PMapPtrTest() string {
	return `
func TestPmap<FTYPE>Ptr(t *testing.T) {
	// Test : square the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v6 <TYPE> = 6
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8
	var v9 <TYPE> = 9

	expectedSquareList := []*<TYPE>{&v1, &v4, &v9}
	squareList := PMap<FTYPE>Ptr(square<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMap<FTYPE>Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*<TYPE>{&v6, &v7, &v8}
	partialAdd<FTYPE>Ptr := func(num *<TYPE>) *<TYPE> {
		r := 5 + *num
		return &r
	}
	sumList := PMap<FTYPE>Ptr(partialAdd<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMap<FTYPE>Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMap<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("PMap<FTYPE>Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMap<FTYPE>Ptr(nil, []*<TYPE>{})) > 0 {
		t.Errorf("PMap<FTYPE>Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func square<FTYPE>Ptr(num *<TYPE>) *<TYPE> {
	r := *num * *num
	return &r
}

`
}

// PMap<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func PMapPtrBoolTest() string {
	return `
func TestPMap<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []*<TYPE>{&vf}
	
	newList := PMap<FTYPE>Ptr(inverse<FTYPE>Ptr, []*<TYPE>{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("Map<FTYPE>Ptr failed")
	}

	if len(PMap<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FTYPE>Ptr failed.")
	}
}

`
}

// ************PMapPtrErr*************************
// PMap<FTYPE>PtrErr applies the function(1st argument) on each item of the list and returns new list
func PMapPtrErrTest() string {
	return `
func TestPmap<FTYPE>PtrErr(t *testing.T) {
	// Test : square the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v6 <TYPE> = 6
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8
	var v9 <TYPE> = 9
	var v0 <TYPE> = 0

	expectedSquareList := []*<TYPE>{&v1, &v4, &v9}
	squareList, _ := PMap<FTYPE>PtrErr(square<FTYPE>PtrErr, []*<TYPE>{&v1, &v2, &v3})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMap<FTYPE>Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*<TYPE>{&v6, &v7, &v8}
	partialAdd<FTYPE>PtrErr := func(num *<TYPE>) (*<TYPE>, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMap<FTYPE>PtrErr(partialAdd<FTYPE>PtrErr, []*<TYPE>{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMap<FTYPE>PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMap<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FTYPE>PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMap<FTYPE>PtrErr(nil, []*<TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FTYPE>PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMap<FTYPE>PtrErr(square<FTYPE>PtrErr, []*<TYPE>{&v1, &v0, &v0, &v0, &v3})
	if err == nil {
		t.Errorf("PMap<FTYPE>PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func square<FTYPE>PtrErr(num *<TYPE>) (*<TYPE>, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

`
}

// PMap<FTYPE>PtrErr applies the function(1st argument) on each item of the list and returns new list
func PMapPtrErrBoolTest() string {
	return `
func TestPMap<FTYPE>PtrErr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []*<TYPE>{&vf}
	
	newList, _ := PMap<FTYPE>PtrErr(inverse<FTYPE>PtrErr, []*<TYPE>{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("Map<FTYPE>PtrErr failed")
	}

	r, _ := PMap<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FTYPE>PtrErr failed.")
	}

	_, err := PMap<FTYPE>PtrErr(inverse<FTYPE>PtrErr, []*<TYPE>{&vf, &vf, &vf, &vt})
	if err == nil {
		t.Errorf("Map<FTYPE>PtrErr failed")
	}
}

`
}

func ReplaceActivityPMapPtrErr(code string) string {
	s1 := `import (
    _ "errors"
	"reflect"
	"testing"
)

func TestPmapIntPtrErr(t *testing.T) {`
	s2 := `import (
    "errors"
	"testing"
)

func TestPmapIntPtrErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestPmapStrPtrErr(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v6 string = "6"
	var v7 string = "7"
	var v8 string = "8"
	var v9 string = "9"
	var v0 string = "0"

	expectedSquareList := []*string{&v1, &v4, &v9}
	squareList, _ := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapStrPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}
	partialAddStrPtrErr := func(num *string) (*string, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapStrPtrErr(partialAddStrPtrErr, []*string{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapStrPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v0, &v0, &v0, &v3})
	if err == nil {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareStrPtrErr(num *string) (*string, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}`
	s2 = `func TestPmapStrPtrErr(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v0 string = "0"

	var v51 string = "51"
	var v52 string = "52"
	var v53 string = "53"

	var v11 string = "11"
	var v22 string = "22"
	var v33 string = "33"

	expectedSquareList := []*string{&v11, &v22, &v33}
	squareList, _ := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapStrPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v51, &v52, &v53}
	partialAddStrPtrErr := func(num *string) (*string, error) {
		if *num == "0" {
			return nil, errors.New("0 is invalid for this test")
		}
		r := "5" + *num
		return &r, nil
	}
	sumList, _ := PMapStrPtrErr(partialAddStrPtrErr, []*string{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapStrPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v0, &v0, &v0, &v0, &v3})
	if err == nil {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareStrPtrErr(num *string) (*string, error) {
	if *num == "0" {
		return nil, errors.New("0  is invalid number")
	}
	r := *num + *num
	return &r, nil
}`
	code = strings.Replace(code, s1, s2, -1)
	return code
}

// ************PMapErr*************************
// PMap<FTYPE>Err applies the function(1st argument) on each item of the list and returns new list
func PMapErrTest() string {
	return `
func TestPmap<FTYPE>Err(t *testing.T) {
	// Test : square the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v6 <TYPE> = 6
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8
	var v9 <TYPE> = 9
	var v0 <TYPE> = 0

	expectedSquareList := []<TYPE>{v1, v4, v9}
	squareList, _ := PMap<FTYPE>Err(square<FTYPE>Err, []<TYPE>{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMap<FTYPE>Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []<TYPE>{v6, v7, v8}
	partialAdd<FTYPE>Err := func(num <TYPE>) (<TYPE>, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMap<FTYPE>Err(partialAdd<FTYPE>Err, []<TYPE>{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMap<FTYPE>Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMap<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMap<FTYPE>Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMap<FTYPE>Err(nil, []<TYPE>{})
	if len(r) > 0 {
		t.Errorf("PMap<FTYPE>Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMap<FTYPE>Err(square<FTYPE>Err, []<TYPE>{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMap<FTYPE>Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func square<FTYPE>Err(num <TYPE>) (<TYPE>, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}

`
}

// PMap<FTYPE>Err applies the function(1st argument) on each item of the list and returns new list
func PMapErrBoolTest() string {
	return `
func TestPMap<FTYPE>Err(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []<TYPE>{vf}
	
	newList, _ := PMap<FTYPE>Err(inverse<FTYPE>Err, []<TYPE>{vt})
	if newList[0] != expectedSumList[0]  {
		t.Errorf("Map<FTYPE>Err failed")
	}

	r, _ := PMap<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FTYPE>Err failed.")
	}

	_, err := PMap<FTYPE>Err(inverse<FTYPE>Err, []<TYPE>{vf, vf, vf, vt})
	if err == nil {
		t.Errorf("Map<FTYPE>Err failed")
	}
}

`
}

func ReplaceActivityPMapErr(code string) string {
	s1 := `import (
    _ "errors"
	"reflect"
	"testing"
)

func TestPmapIntErr(t *testing.T) {`
	s2 := `import (
    "errors"
	"testing"
)

func TestPmapIntErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestPmapStrErr(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v6 string = "6"
	var v7 string = "7"
	var v8 string = "8"
	var v9 string = "9"
	var v0 string = "0"

	expectedSquareList := []string{v1, v4, v9}
	squareList, _ := PMapStrErr(squareStrErr, []string{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapStrErr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []string{v6, v7, v8}
	partialAddStrErr := func(num string) (string, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapStrErr(partialAddStrErr, []string{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapStrErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapStrErr(squareStrErr, []string{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareStrErr(num string) (string, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}`
	s2 = `func TestPmapStrErr(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v0 string = "0"

	var v51 string = "51"
	var v52 string = "52"
	var v53 string = "53"

	var v11 string = "11"
	var v22 string = "22"
	var v33 string = "33"

	expectedSquareList := []string{v11, v22, v33}
	squareList, _ := PMapStrErr(squareStrErr, []string{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapStrErr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []string{v51, v52, v53}
	partialAddStrErr := func(num string) (string, error) {
		if num == "0" {
			return "", errors.New("0 is invalid for this test")
		}
		r := "5" + num
		return r, nil
	}
	sumList, _ := PMapStrErr(partialAddStrErr, []string{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapStrErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapStrErr(squareStrErr, []string{v1, v0, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareStrErr(num string) (string, error) {
	if num == "0" {
		return "0", errors.New("0  is invalid number")
	}
	r := num + num
	return r, nil
}`
	code = strings.Replace(code, s1, s2, -1)
	return code
}
