package basic

import "strings"

// FilterPtrTest applies the function(1st argument) on each item of the list and returns new list
func FilterPtrTest() string {
	return `
func TestFilter<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v10 <TYPE> = 10
	var v20 <TYPE> = 20
	var v40 <TYPE> = 40


	// Test : even number in the list
	expectedFilteredList := []*<TYPE>{&v2, &v4}
	filteredList := Filter<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*<TYPE>{&v20, &v40}
	partialIsEven := func(num *<TYPE>) bool { return isEvenDivisibleBy<FTYPE>Ptr(num, &v10) }
	filteredList = Filter<FTYPE>Ptr(partialIsEven, []*<TYPE>{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(Filter<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEven<FTYPE>Ptr(num *<TYPE>) bool {
	return *num%2 == 0
}

func isEvenDivisibleBy<FTYPE>Ptr(num, divisibleBy *<TYPE>) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}

`
}

// FilterPtrBoolTest applies the function(1st argument) on each item of the list and returns new list
func FilterPtrBoolTest() string {
	return `
func TestFilter<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true

	expectedSumList := []*<TYPE>{&vt}
	
	newList := Filter<FTYPE>Ptr(true<FTYPE>Ptr, []*<TYPE>{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("Filter<FTYPE>Ptr failed")
	}

	if len(Filter<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Map<FTYPE>Ptr failed.")
	}
}

func true<FTYPE>Ptr(num1 *<TYPE>) bool {
	return true
}

`
}

// FilterPtrErrTest applies the function(1st argument) on each item of the list and returns new list
func FilterPtrErrTest() string {
	return `
func TestFilter<FTYPE>PtrErr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v0 <TYPE> = 0

	// Test : even number in the list
	expectedFilteredList := []*<TYPE>{&v2, &v4}
	filteredList, _ := Filter<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := Filter<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Filter<FTYPE>PtrErr failed.")
	}

	_, err := Filter<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{&v0})
	if err == nil {
		t.Errorf("Filter<FTYPE>PtrErr failed.")
	}
}

func isEven<FTYPE>PtrErr(num *<TYPE>) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}

`
}

// FilterPtrErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func FilterPtrErrBoolTest() string {
	return `
func TestFilter<FTYPE>PtrErr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []*<TYPE>{&vt}
	
	newList, _ := Filter<FTYPE>PtrErr(true<FTYPE>PtrErr, []*<TYPE>{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("Filter<FTYPE>PtrErr failed")
	}

	r, _ := Filter<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Filter<FTYPE>PtrErr failed.")
	}

	_, err := Filter<FTYPE>PtrErr(true<FTYPE>PtrErr, []*<TYPE>{&vf})
	if err == nil {
		t.Errorf("Filter<FTYPE>PtrErr failed.")
	}
}

func true<FTYPE>PtrErr(num1 *<TYPE>) (bool, error) {
	if *num1 == false {
		return false, errors.New("False is not allowed")
	}
	return true, nil
}

`
}

// ReplaceActivityFilterPtrErrTest is template
func ReplaceActivityFilterPtrErrTest(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestFilterIntPtrErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestFilterIntPtrErr(t *testing.T) {`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isEvenStrPtrErr(num *string) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}`

	s2 = `func isEvenStrPtrErr(num *string) (bool, error) {
	if *num == "0" {
		return false, errors.New("Zero is not allowed")
	} else if *num == "2" || *num == "4" {
		return true, nil
	}
	return false, nil
	
}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isEvenFloat32PtrErr(num *float32) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}`

	s2 = `func isEvenFloat32PtrErr(num *float32) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return int(*num)%2 == 0, nil
}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isEvenFloat64PtrErr(num *float64) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}`

	s2 = `func isEvenFloat64PtrErr(num *float64) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return int(*num)%2 == 0, nil
}`

	code = strings.Replace(code, s1, s2, -1)

	return code
}

// FilterErrTest applies the function(1st argument) on each item of the list and returns new list
func FilterErrTest() string {
	return `
func TestFilter<FTYPE>Err(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v0 <TYPE> = 0

	// Test : even number in the list
	expectedFilteredList := []<TYPE>{v2, v4}
	filteredList, _ := Filter<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := Filter<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Filter<FTYPE>Err failed.")
	}

	_, err := Filter<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{v0})
	if err == nil {
		t.Errorf("Filter<FTYPE>PtrErr failed.")
	}
}

func isEven<FTYPE>Err(num <TYPE>) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}
`
}

// FilterErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func FilterErrBoolTest() string {
	return `
func TestFilter<FTYPE>Err(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedSumList := []<TYPE>{vt}

	newList, _ := Filter<FTYPE>Err(true<FTYPE>Err, []<TYPE>{vt})
	if newList[0] != expectedSumList[0] {
		t.Errorf("Filter<FTYPE>Err failed")
	}

	r, _ := Filter<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Filter<FTYPE>Err failed.")
	}

	_, err := Filter<FTYPE>Err(true<FTYPE>Err, []<TYPE>{vf})
	if err == nil {
		t.Errorf("Filter<FTYPE>Err failed.")
	}
}

func true<FTYPE>Err(num1 <TYPE>) (bool, error) {
	if num1 == false {
		return false, errors.New("False is not allowed")
	}
	return true, nil
}
`
}

// ReplaceActivityFilterErrTest replaces ...
func ReplaceActivityFilterErrTest(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestFilterIntErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestFilterIntErr(t *testing.T) {`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isEvenStrErr(num string) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}`

	s2 = `func isEvenStrErr(num string) (bool, error) {
	if num == "0" {
		return false, errors.New("Zero is not allowed")
	} else if num == "2" || num == "4" {
		return true, nil
	}
	return false, nil

}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isEvenFloat32Err(num float32) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}`

	s2 = `func isEvenFloat32Err(num float32) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return int(num)%2 == 0, nil
}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isEvenFloat64Err(num float64) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}`

	s2 = `func isEvenFloat64Err(num float64) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return int(num)%2 == 0, nil
}`

	code = strings.Replace(code, s1, s2, -1)

	return code
}
