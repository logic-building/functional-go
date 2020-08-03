package basic

import "strings"

// FilterMapPtrTest is template
func FilterMapPtrTest() string {
	return `
func TestFilterMap<FTYPE>Ptr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 <TYPE> = 1
	var v4 <TYPE> = 4
	var v8 <TYPE> = 8
	var v0 <TYPE>
	var v2 <TYPE> = 2

	expectedFilteredList := []*<TYPE>{&v2, &v4, &v8}
	filteredList := FilterMap<FTYPE>Ptr(isPositive<FTYPE>Ptr, multiplyBy2<FTYPE>Ptr, []*<TYPE>{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMap<FTYPE>Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMap<FTYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositive<FTYPE>Ptr(num *<TYPE>) bool {
	return *num > 0
}
func multiplyBy2<FTYPE>Ptr(num *<TYPE>) *<TYPE> {
	result := *num * 2
	return &result
}
`
}

// FilterMapPtrBoolTest is template
func FilterMapPtrBoolTest() string {
	return `
func TestFilterMap<FTYPE>Ptr(t *testing.T) {
	var vt1 <TYPE> = true
	var vt2 <TYPE> = true
	var vf1 <TYPE> = false
	var vf2 <TYPE> = false

	expectedFilteredList := []*<TYPE>{&vt1, &vt2}
	filteredList := FilterMap<FTYPE>Ptr(isTrue<FTYPE>Ptr, returnSame<FTYPE>Ptr, []*<TYPE>{&vt1, &vt2, &vf1, &vf2})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMap<FTYPE>Ptr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMap<FTYPE>Ptr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMap<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isTrue<FTYPE>Ptr(num *<TYPE>) bool {
	return *num == true
}
func returnSame<FTYPE>Ptr(num *<TYPE>) *<TYPE> {
	return num
}
`
}

// FilterMapPtrErrTest is template
func FilterMapPtrErrTest() string {
	return `
func TestFilterMap<FTYPE>PtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 <TYPE> = 1
	var v4 <TYPE> = 4
	var v8 <TYPE> = 8
	var v16 <TYPE> = 16
	var v0 <TYPE>
	var v2 <TYPE> = 2

	expectedFilteredList := []*<TYPE>{&v8, &v16}
	filteredList, _ := FilterMap<FTYPE>PtrErr(isPositive<FTYPE>PtrErr, multiplyBy2<FTYPE>PtrErr, []*<TYPE>{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMap<FTYPE>PtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMap<FTYPE>PtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FTYPE>Ptr failed.")
	}

	_, err := FilterMap<FTYPE>PtrErr(isPositive<FTYPE>PtrErr, multiplyBy2<FTYPE>PtrErr, []*<TYPE>{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMap<FTYPE>PtrErr failed.")
	}

	_, err = FilterMap<FTYPE>PtrErr(isPositive<FTYPE>PtrErr, multiplyBy2<FTYPE>PtrErr, []*<TYPE>{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMap<FTYPE>PtrErr failed.")
	}
}

func isPositive<FTYPE>PtrErr(num *<TYPE>) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil
}

func multiplyBy2<FTYPE>PtrErr(num *<TYPE>) (*<TYPE>, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil
}
`
}

// ReplaceActivityFilterMapPtrErr replaces ...
func ReplaceActivityFilterMapPtrErr(code string) string {
	s1 := `_ "errors"
	"reflect"
	"testing"
)

func TestFilterMapIntPtrErr(t *testing.T) {`

	s2 := `"errors"
	"testing"
)

func TestFilterMapIntPtrErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isPositiveStrPtrErr(num *string) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return *num > 1, nil`
	s2 = `func isPositiveStrPtrErr(num *string) (bool, error) {
	if *num == "0" {
		return false, errors.New("Zero is invalid")
	}
	if *num == "2" || *num == "3" || *num == "4" || *num == "5" || *num == "6" || *num == "7" || *num == "8" {
		return true, nil
	} else {
		return false, nil
	}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func multiplyBy2StrPtrErr(num *string) (*string, error) {
	if *num == 2 {
		return nil, errors.New("2 is invalid number")
	}
	result := *num * 2
	return &result, nil`

	s2 = `func multiplyBy2StrPtrErr(num *string) (*string, error) {
	result := "100"
	if *num == "2" {
		return nil, errors.New("2 is invalid number")
	}
	if *num == "3" {
		result = "6"
	}

	if *num == "4" {
		result = "8"
	}

	if *num == "5" {
		result = "10"
	}

	if *num == "6" {
		result = "12"
	}

	if *num == "7" {
		result = "14"
	}

	if *num == "8" {
		result = "16"
	}

	return &result, nil`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestFilterMapStrPtr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 string = "1"
	var v4 string = "4"
	var v8 string = "8"
	var v0 string
	var v2 string = "2"

	expectedFilteredList := []*string{&v2, &v4, &v8}
	filteredList := FilterMapStrPtr(isPositiveStrPtr, multiplyBy2StrPtr, []*string{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2] {
		t.Errorf("FilterMapStrPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapStrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveStrPtr(num *string) bool {
	return *num > 0
}
func multiplyBy2StrPtr(num *string) *string {
	result := *num * 2
	return &result
}`
	s2 = `func TestFilterMapStrPtr(t *testing.T) {
	var v1 string = "1"
	var v4 string = "4"
	var v0 string = "0"
	var v2 string = "2"

	expectedFilteredList := []*string{&v1, &v0, &v2, &v4}
	filteredList := FilterMapStrPtr(noFilter, concatA, []*string{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0]+"A" || *filteredList[1] != *expectedFilteredList[1]+"A" {
		t.Errorf("FilterMapStrPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapStrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrPtr failed.")
	}
}

func noFilter(num *string) bool {
	return true
}
func concatA(num *string) *string {
	result := *num + "A"
	return &result
}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestFilterMapStrPtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 string = "1"
	var v4 string = "4"
	var v8 string = "8"
	var v16 string = "16"
	var v0 string
	var v2 string = "2"

	expectedFilteredList := []*string{&v8, &v16}
	filteredList, _ := FilterMapStrPtrErr(isPositiveStrPtrErr, multiplyBy2StrPtrErr, []*string{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapStrPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapStrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapStrPtr failed.")
	}

	_, err := FilterMapStrPtrErr(isPositiveStrPtrErr, multiplyBy2StrPtrErr, []*string{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapStrPtrErr failed.")
	}

	_, err = FilterMapStrPtrErr(isPositiveStrPtrErr, multiplyBy2StrPtrErr, []*string{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapStrPtrErr failed.")
	}
}`

	s2 = `func TestFilterMapStrPtrErr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 string = "1"
	var v4 string = "4"
	var v8 string = "8"
	var v16 string = "16"
	var v0 string = "0"
	var v2 string = "2"

	expectedFilteredList := []*string{&v8, &v16}
	filteredList, _ := FilterMapStrPtrErr(isPositiveStrPtrErr, multiplyBy2StrPtrErr, []*string{&v1, &v4, &v8})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("FilterMapStrPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMapStrPtrErr(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMapStrPtr failed.")
	}

	_, err := FilterMapStrPtrErr(isPositiveStrPtrErr, multiplyBy2StrPtrErr, []*string{&v0, &v1})
	if err == nil {
		t.Errorf("FilterMapStrPtrErr failed.")
	}

	_, err = FilterMapStrPtrErr(isPositiveStrPtrErr, multiplyBy2StrPtrErr, []*string{&v1, &v2})
	if err == nil {
		t.Errorf("FilterMapStrPtrErr failed.")
	}
}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}
