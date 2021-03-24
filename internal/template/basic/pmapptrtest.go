package basic

import "strings"

// PMap2Test applies the function(1st argument) on each item of the list and returns new list
func PMap2Test() string {
	return `
func TestPmap<FTYPE>2(t *testing.T) {
	// Test : square the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v9 <TYPE> = 9

	expectedSquareList := []<TYPE>{v1, v4, v9}
	squareList := PMap<FTYPE>(square<FTYPE>, []<TYPE>{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMap<FTYPE> failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}
`
}

// PMap2BoolTest applies the function(1st argument) on each item of the list and returns new list
func PMap2BoolTest() string {
	return `
func TestPmap<FTYPE>2(t *testing.T) {
	// Test : square the list
	var vt <TYPE> = true
	var vf <TYPE> = false

	inverse := func(v bool) bool {
		if v == true {
			return false
		}
		return true
	}

	expectedList := []<TYPE>{vt, vt, vt}
	actualList := PMap<FTYPE>(inverse, []<TYPE>{vf, vf, vf}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedList {
		for _, x := range actualList {
			if v == x {
				count++
				break
			}
		}
	}

	if count != len(expectedList) {
		t.Errorf("PMap<FTYPE> failed.expected len=%v, actual len=%v", len(expectedList), count)
		t.Errorf(reflect.String.String())
	}
}
`
}

// ReplaceActivityPMap2Test replaces ...
func ReplaceActivityPMap2Test(code string) string {
	s1 := `func TestPmapStr2(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v9 string = "9"

	expectedSquareList := []string{v1, v4, v9}
	squareList := PMapStr(squareStr, []string{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}`
	s2 := `func TestPmapStr2(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v11 string = "33"
	var v4 string = "22"
	var v9 string = "33"

	appendStr := func(v string) string {
		r := v + v
		return r
	}

	expectedSquareList := []string{v11, v4, v9}
	squareList := PMapStr(appendStr, []string{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}

// PMapPtrTest applies the function(1st argument) on each item of the list and returns new list
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
	squareList := PMap<FTYPE>Ptr(square<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMap<FTYPE>Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*<TYPE>{&v6, &v7, &v8}
	partialAdd<FTYPE>Ptr := func(num *<TYPE>) *<TYPE> {
		r := 5 + *num
		return &r
	}
	sumList := PMap<FTYPE>Ptr(partialAdd<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3}, Optional{FixedPool: 1})
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

// PMapPtr2Test applies the function(1st argument) on each item of the list and returns new list
func PMapPtr2Test() string {
	return `
func TestPmap<FTYPE>Ptr2(t *testing.T) {
	// Test : square the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v9 <TYPE> = 9

	expectedSquareList := []*<TYPE>{&v1, &v4, &v9}
	squareList := PMap<FTYPE>Ptr(square<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMap<FTYPE>Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}
`
}

// PMap2PtrBoolTest applies the function(1st argument) on each item of the list and returns new list
func PMap2PtrBoolTest() string {
	return `
func TestPmap<FTYPE>2Ptr(t *testing.T) {
	// Test : square the list
	var vt <TYPE> = true
	var vf <TYPE> = false

	inverse := func(v *bool) *bool {
		if *v == true {
			return &vf
		}
		return &vt
	}

	expectedList := []*<TYPE>{&vt, &vt, &vt}
	actualList := PMap<FTYPE>Ptr(inverse, []*<TYPE>{&vf, &vf, &vf}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedList {
		for _, x := range actualList {
			if *v == *x {
				count++
				break
			}
		}
	}

	if count != len(expectedList) {
		t.Errorf("PMap<FTYPE> failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}
`
}

// ReplaceActivityPMapPtr2Test replaces ...
func ReplaceActivityPMapPtr2Test(code string) string {
	s1 := `func TestPmapStrPtr2(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v9 string = "9"

	expectedSquareList := []*string{&v1, &v4, &v9}
	squareList := PMapStrPtr(squareStrPtr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStrPtr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}`
	s2 := `func TestPmapStrPtr2(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v11 string = "33"
	var v4 string = "22"
	var v9 string = "33"

	appendStr := func(v *string) *string {
		r := *v + *v
		return &r
	}

	expectedSquareList := []*string{&v11, &v4, &v9}
	squareList := PMapStrPtr(appendStr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStrPtr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}

// PMapPtrBoolTest applies the function(1st argument) on each item of the list and returns new list
func PMapPtrBoolTest() string {
	return `
func TestPMap<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []*<TYPE>{&vf, &vf}

	newList := PMap<FTYPE>Ptr(inverse<FTYPE>Ptr, []*<TYPE>{&vt, &vt}, Optional{FixedPool: 1})
	if *newList[0] != *expectedSumList[0] || *newList[1] != *expectedSumList[1] {
		t.Errorf("Map<FTYPE>Ptr failed")
	}

	if len(PMap<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FTYPE>Ptr failed.")
	}
}
`
}

func ReplaceActivityMapPtr(code string) string {
	s1 := `func TestPmapStrPtr(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v6 string = "6"
	var v7 string = "7"
	var v8 string = "8"
	var v9 string = "9"

	expectedSquareList := []*string{&v1, &v4, &v9}
	squareList := PMapStrPtr(squareStrPtr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapStrPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}
	partialAddStrPtr := func(num *string) *string {
		r := 5 + *num
		return &r
	}
	sumList := PMapStrPtr(partialAddStrPtr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapStrPtr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapStrPtr(nil, nil)) > 0 {
		t.Errorf("PMapStrPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapStrPtr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareStrPtr(num *string) *string {
	r := *num * *num
	return &r
}`

	s2 := `func TestPmapStrPtr(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v6 string = "51"
	var v7 string = "52"
	var v8 string = "53"

	var v1SquareExpected string = "11"
	var v2SquareExpected string = "22"
	var v3SquareExpected string = "33"

	expectedSquareList := []*string{&v1SquareExpected, &v2SquareExpected, &v3SquareExpected}
	squareList := PMapStrPtr(squareStrPtr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapStrPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}
	partialAddStrPtr := func(num *string) *string {
		r := "5" + *num
		return &r
	}
	sumList := PMapStrPtr(partialAddStrPtr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapStrPtr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapStrPtr(nil, nil)) > 0 {
		t.Errorf("PMapStrPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapStrPtr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareStrPtr(num *string) *string {
	r := *num + *num
	return &r
}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}

// ************PMapPtrErr*************************

// PMapPtrErrTest applies the function(1st argument) on each item of the list and returns new list
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
	var v0 <TYPE>

	expectedSquareList := []*<TYPE>{&v1, &v4, &v9}
	squareList, _ := PMap<FTYPE>PtrErr(square<FTYPE>PtrErr, []*<TYPE>{&v1, &v2, &v3}, Optional{FixedPool: 1})

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
	sumList, _ := PMap<FTYPE>PtrErr(partialAdd<FTYPE>PtrErr, []*<TYPE>{&v1, &v2, &v3}, Optional{FixedPool: 1})
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

	_, err := PMap<FTYPE>PtrErr(square<FTYPE>PtrErr, []*<TYPE>{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMap<FTYPE>PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMap<FTYPE>PtrErr(square<FTYPE>PtrErr, []*<TYPE>{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMap<FTYPE>PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMap<FTYPE>PtrErr(square<FTYPE>PtrErr, []*<TYPE>{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
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

// PMapPtrErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func PMapPtrErrBoolTest() string {
	return `
func TestPMap<FTYPE>PtrErr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []*<TYPE>{&vf, &vf}

	newList, _ := PMap<FTYPE>PtrErr(inverse<FTYPE>PtrErr, []*<TYPE>{&vt, &vt}, Optional{FixedPool: 1})
	if *newList[0] != *expectedSumList[0] || *newList[1] != *expectedSumList[1] {
		t.Errorf("Map<FTYPE>PtrErr failed")
	}

	r, _ := PMap<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FTYPE>PtrErr failed.")
	}

	_, err := PMap<FTYPE>PtrErr(inverse<FTYPE>PtrErr, []*<TYPE>{&vf, &vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("Map<FTYPE>PtrErr failed")
	}
}
`
}

// ReplaceActivityPMapPtrErr replaces ...
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
	"time"
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
	var v0 string

	expectedSquareList := []*string{&v1, &v4, &v9}
	squareList, _ := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})

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
	sumList, _ := PMapStrPtrErr(partialAddStrPtrErr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})
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

	_, err := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
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
	squareList, _ := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})

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
	sumList, _ := PMapStrPtrErr(partialAddStrPtrErr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})
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

	_, err := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v0, &v0, &v0, &v0, &v3}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
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

// PMapPtr2ErrTest applies the function(1st argument) on each item of the list and returns new list
func PMapPtr2ErrTest() string {
	return `
func TestPmap<FTYPE>Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v9 <TYPE> = 9
	var v0 <TYPE>

	expectedSquareList := []*<TYPE>{&v1, &v4, &v9}
	squareList, _ := PMap<FTYPE>PtrErr(square<FTYPE>PtrErr, []*<TYPE>{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMap<FTYPE>Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMap<FTYPE>PtrErr(square<FTYPE>PtrErr, []*<TYPE>{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMap<FTYPE>PtrErr(square<FTYPE>PtrErr, []*<TYPE>{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMap<FTYPE>PtrErr(square2<FTYPE>PtrErr, []*<TYPE>{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMap<FTYPE>PtrErr(square2<FTYPE>PtrErr, []*<TYPE>{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMap<FTYPE>PtrErr(square2<FTYPE>PtrErr, []*<TYPE>{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2<FTYPE>PtrErr(num *<TYPE>) (*<TYPE>, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}
`
}

// PMap2PtrErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func PMapPtr2ErrBoolTest() string {
	return `
func TestPmap<FTYPE>2PtrErr(t *testing.T) {
	// Test : square the list
	var vt <TYPE> = true
	var vf <TYPE> = false

	inverse := func(v *bool) (*bool, error) {
		if *v == true {
			return &vt, nil
		}

		if *v == false {
			return nil, errors.New("false is error")
		}
		return &vf, nil
	}

	inverse2 := func(v *bool) (*bool, error) {
		time.Sleep(50 * time.Millisecond)
		if *v == true {
			return &vt, nil
		}

		if *v == false {
			return nil, errors.New("false is error")
		}
		return &vf, nil
	}

	expectedList := []*<TYPE>{&vt, &vt, &vt}
	actualList, _ := PMap<FTYPE>PtrErr(inverse, []*<TYPE>{&vt, &vt, &vt}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedList {
		for _, x := range actualList {
			if *v == *x {
				count++
				break
			}
		}
	}

	if count != len(expectedList) {
		t.Errorf("PMap<FTYPE>Ptr2Err failed.expected len=%v, actual len=%v", len(expectedList), count)
	}

	_, err := PMap<FTYPE>PtrErr(inverse, []*<TYPE>{&vf, &vf, &vt}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed")
	}

	_, err = PMap<FTYPE>PtrErr(inverse, []*<TYPE>{&vf, &vf, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed")
	}

	_, err = PMap<FTYPE>PtrErr(inverse, []*<TYPE>{&vf, &vf, &vf, &vf, &vf, &vt}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed")
	}

	_, err = PMap<FTYPE>PtrErr(inverse2, []*<TYPE>{&vf, &vt}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed")
	}

	_, err = PMap<FTYPE>PtrErr(inverse, []*<TYPE>{&vf, &vf, &vf, &vf, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed")
	}
}
`
}

// ReplaceActivityPMapPtr2ErrTest replaces ...
func ReplaceActivityPMapPtr2ErrTest(code string) string {
	s1 := `func TestPmapStrPtr2Err(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v9 string = "9"
	var v0 string

	expectedSquareList := []*string{&v1, &v4, &v9}
	squareList, _ := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapStrPtrErr(squareStrPtrErr, []*string{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(squareStrPtrErr, []*string{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(square2StrPtrErr, []*string{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(square2StrPtrErr, []*string{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(square2StrPtrErr, []*string{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2StrPtrErr(num *string) (*string, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}`
	s2 := `func TestPmapStrPtr2Err(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v11 string = "33"
	var v4 string = "22"
	var v9 string = "33"
	var v0 string = "0"

	appendStr := func(v *string) (*string, error) {
		r := *v + *v
		if *v == "0" {
			return nil, errors.New("0 is not allowed in this test")
		}
		return &r, nil
	}

	appendStr2 := func(v *string) (*string, error) {
		time.Sleep(200 * time.Millisecond)
		r := *v + *v
		if *v == "0" {
			return nil, errors.New("0 is not allowed in this test")
		}
		return &r, nil
	}

	expectedSquareList := []*string{&v11, &v4, &v9}
	squareList, _ := PMapStrPtrErr(appendStr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapStrPtrErr(squareStrPtrErr, []*string{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(squareStrPtrErr, []*string{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(appendStr2, []*string{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(appendStr2, []*string{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(appendStr, []*string{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}

// ************PMapErr*************************

// PMapErrTest applies the function(1st argument) on each item of the list and returns new list
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
	var v0 <TYPE>

	expectedSquareList := []<TYPE>{v1, v4, v9}
	squareList, _ := PMap<FTYPE>Err(square<FTYPE>Err, []<TYPE>{v1, v2, v3}, Optional{FixedPool: 2})

	for _, v := range squareList {
		if !Exists<FTYPE>(v, expectedSquareList) {
			t.Errorf("PMap<FType>Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
		}
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
	sumList, _ := PMap<FTYPE>Err(partialAdd<FTYPE>Err, []<TYPE>{v1, v2, v3}, Optional{FixedPool: 2})
	for _, v := range sumList {
		if !Exists<FTYPE>(v, expectedSumList) {
			t.Errorf("PMap<FTYPE>Err failed.expected=%v, actual=%v", expectedSumList, sumList)
		}
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

	_, err = PMap<FTYPE>Err(square<FTYPE>Err, []<TYPE>{v0, v0, v0, v0, v3}, Optional{FixedPool: 2})
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

// PMapErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func PMapErrBoolTest() string {
	return `
func TestPMap<FTYPE>Err(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []<TYPE>{vf}

	newList, _ := PMap<FTYPE>Err(inverse<FTYPE>Err, []<TYPE>{vt})
	if newList[0] != expectedSumList[0] {
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

// ReplaceActivityPMapErr replaces ...
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
	"time"
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
	var v0 string

	expectedSquareList := []string{v1, v4, v9}
	squareList, _ := PMapStrErr(squareStrErr, []string{v1, v2, v3}, Optional{FixedPool: 2})

	for _, v := range squareList {
		if !ExistsStr(v, expectedSquareList) {
			t.Errorf("PMap<FType>Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
		}
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
	sumList, _ := PMapStrErr(partialAddStrErr, []string{v1, v2, v3}, Optional{FixedPool: 2})
	for _, v := range sumList {
		if !ExistsStr(v, expectedSumList) {
			t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
		}
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

	_, err = PMapStrErr(squareStrErr, []string{v0, v0, v0, v0, v3}, Optional{FixedPool: 2})
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
	squareList, _ := PMapStrErr(squareStrErr, []string{v1, v2, v3}, Optional{FixedPool: 2})

	for _, v := range squareList {
		if !ExistsStr(v, expectedSquareList) {
			t.Errorf("PMap<FType>Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
		}
	}

	// Test: add 5 to each item in the list
	expectedSumList := []string{v51, v52, v53}
	partialAddStrErr := func(num string) (string, error) {
		if num == "0" {
			return "0", errors.New("0 is invalid for this test")
		}
		r := "5" + num
		return r, nil
	}
	sumList, _ := PMapStrErr(partialAddStrErr, []string{v1, v2, v3}, Optional{FixedPool: 2})
	for _, v := range sumList {
		if !ExistsStr(v, expectedSumList) {
			t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
		}
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

	_, err = PMapStrErr(squareStrErr, []string{v0, v0, v0, v0, v3}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrErr failed.expected=error")
	}

	_, err = PMapStrErr(squareStrErr, []string{v0, v0, v0, v0, v3}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrErr failed.expected=error")
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

// PMap2ErrTest applies the function(1st argument) on each item of the list and returns new list
func PMap2ErrTest() string {
	return `
func TestPmap<FTYPE>2Err(t *testing.T) {
	// Test : square the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v9 <TYPE> = 9
	var v0 <TYPE>

	expectedSquareList := []<TYPE>{v1, v4, v9}
	squareList, _ := PMap<FTYPE>Err(square<FTYPE>Err, []<TYPE>{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMap<FTYPE>2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMap<FTYPE>Err(square<FTYPE>Err, []<TYPE>{v0, v0, v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMap<FTYPE>Err(square2<FTYPE>Err, []<TYPE>{v0, v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMap<FTYPE>Err(square2<FTYPE>Err, []<TYPE>{v0, v3})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMap<FTYPE>Err(square2<FTYPE>Err, []<TYPE>{v3, v0})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>2Err failed")
	}
}

func square2<FTYPE>Err(num <TYPE>) (<TYPE>, error) {
	time.Sleep(50 * time.Millisecond)
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}
`
}

// PMap2ErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func PMap2ErrBoolTest() string {
	return `
func TestPmap<FTYPE>2Err(t *testing.T) {
	// Test : square the list
	var vt <TYPE> = true
	var vf <TYPE> = false

	inverse := func(v bool) (bool, error) {
		if v == true {
			return vt, nil
		}

		if v == false {
			return false, errors.New("false is invalid for this test")
		}
		return vf, nil
	}

	expectedList := []<TYPE>{vt, vt, vt}
	actualList, _ := PMap<FTYPE>Err(inverse, []<TYPE>{vt, vt, vt}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedList {
		for _, x := range actualList {
			if v == x {
				count++
				break
			}
		}
	}

	if count != len(expectedList) {
		t.Errorf("PMap<FTYPE>2Err failed.expected len=%v, actual len=%v", len(expectedList), count)
	}

	_, err := PMap<FTYPE>Err(inverse, []<TYPE>{vf, vf, vt}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FTYPE>2Err failed")
	}

	_, err = PMap<FTYPE>Err(inverse, []<TYPE>{vt, vt, vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMap<FTYPE>2Err failed")
	}

	_, err = PMap<FTYPE>Err(inverse, []<TYPE>{vt, vt, vf}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMap<FTYPE>2Err failed")
	}

	_, err = PMap<FTYPE>Err(inverse, []<TYPE>{vf, vf, vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMap<FTYPE>2Err failed")
	}
}
`
}

// ReplaceActivityPMap2ErrTest replaces ...
func ReplaceActivityPMap2ErrTest(code string) string {
	s1 := `func TestPmapStr2Err(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v9 string = "9"
	var v0 string

	expectedSquareList := []string{v1, v4, v9}
	squareList, _ := PMapStrErr(squareStrErr, []string{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapStrErr(squareStrErr, []string{v0, v0, v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapStr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrErr(square2StrErr, []string{v0, v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapStr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrErr(square2StrErr, []string{v0, v3})
	
	if err == nil {
		t.Errorf("PMapStr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrErr(square2StrErr, []string{v3, v0})
	
	if err == nil {
		t.Errorf("PMapStr2Err failed")
	}
}

func square2StrErr(num string) (string, error) {
	time.Sleep(50 * time.Millisecond)
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}`
	s2 := `func TestPmapStr2Err(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v11 string = "11"
	var v4 string = "22"
	var v9 string = "33"
	var v0 string = "0"

	appendStr := func(v string) (string, error) {
		r := v + v
		if v == "0" {
			return "", errors.New("0 is invalid for this test")
		}
		return r, nil
	}

	appendStr2 := func(v string) (string, error) {
		time.Sleep(50 * time.Millisecond)
		r := v + v
		if v == "0" {
			return "", errors.New("0 is invalid for this test")
		}
		return r, nil
	}

	expectedSquareList := []string{v11, v4, v9}
	squareList, _ := PMapStrErr(appendStr, []string{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapStrErr(appendStr, []string{v0, v0, v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapStr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrErr(appendStr2, []string{v0, v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrErr(appendStr2, []string{v0, v3})
	
	if err == nil {
		t.Errorf("PMapStr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrErr(appendStr2, []string{v3, v0})
	
	if err == nil {
		t.Errorf("PMapStr2Err failed")
	}
}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}
