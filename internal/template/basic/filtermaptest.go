package basic

import (
	"strings"
)

// FilterMapErrTest is template
func FilterMapErrTest() string {
	return `
func TestFilterMap<FTYPE>Err(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 <TYPE> = 1
	var v4 <TYPE> = 4
	var v8 <TYPE> = 8
	var v16 <TYPE> = 16
	var v0 <TYPE> = 0
	var v2 <TYPE> = 2
	
	expectedFilteredList := []<TYPE>{v8, v16}
	filteredList, _ := FilterMap<FTYPE>Err(isPositive<FTYPE>Err, multiplyBy2<FTYPE>Err, []<TYPE>{v1, v4, v8})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FilterMap<FTYPE>Err failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	r, _ := FilterMap<FTYPE>Err(nil, nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterMap<FTYPE>Err failed.")
	}

	_, err := FilterMap<FTYPE>Err(isPositive<FTYPE>Err, multiplyBy2<FTYPE>Err, []<TYPE>{v0, v1})
	if err == nil {
		t.Errorf("FilterMap<FTYPE>Err failed.")
	}

	_, err = FilterMap<FTYPE>Err(isPositive<FTYPE>Err, multiplyBy2<FTYPE>Err, []<TYPE>{v1, v2})
	if err == nil {
		t.Errorf("FilterMap<FTYPE>Err failed.")
	}
}

func isPositive<FTYPE>Err(num <TYPE>) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil
}

func multiplyBy2<FTYPE>Err(num <TYPE>) (<TYPE>, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil
}

`
}

// ReplaceActivityFilterMapErr replaces ...
func ReplaceActivityFilterMapErr(code string) string {
	s1 := `_ "errors"
	"reflect"
	"testing"
)

func TestFilterMapIntErr(t *testing.T) {`

	s2 := `"errors"
	"testing"
)

func TestFilterMapIntErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isPositiveStrErr(num string) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is invalid")
	}
	return num > 1, nil`
	s2 = `func isPositiveStrErr(num string) (bool, error) {
	if num == "0" {
		return false, errors.New("Zero is invalid")
	}
	if num == "2" || num == "3" || num == "4" || num == "5" || num == "6" || num == "7" || num == "8" {
		return true, nil
	}
	return false, nil`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func multiplyBy2StrErr(num string) (string, error) {
	if num == 2 {
		return 0, errors.New("2 is invalid number")
	}
	result := num * 2
	return result, nil`

	s2 = `func multiplyBy2StrErr(num string) (string, error) {
	result := "100"
	if num == "2" {
		return "0", errors.New("2 is invalid number")
	}
	if num == "3" {
		result = "6"
	} 

	if num == "4" {
		result = "8"
	} 

	if num == "5" {
		result = "10"
	} 

	if num == "6" {
		result = "12"
	} 

	if num == "7" {
		result = "14"
	} 

	if num == "8" {
		result = "16"
	} 
	
	return result, nil`

	code = strings.Replace(code, s1, s2, -1)

	return code
}
