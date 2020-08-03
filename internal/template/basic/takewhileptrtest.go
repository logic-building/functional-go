package basic

import "strings"

// TakeWhilePtrTest is template
func TakeWhilePtrTest() string {
	return `
func TestTakeWhile<FTYPE>Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v7 <TYPE> = 7
	var v40 <TYPE> = 40

	expectedNewList := []*<TYPE>{&v4, &v2, &v4}
	NewList := TakeWhile<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*<TYPE>{&v40}
	partialIsEvenDivisibleBy := func(num *<TYPE>) bool { return *num%10 == 0 }
	NewList = TakeWhile<FTYPE>Ptr(partialIsEvenDivisibleBy, []*<TYPE>{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhile<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
	}

	if len(TakeWhile<FTYPE>Ptr(nil, []*<TYPE>{})) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}
`
}

// TakeWhileTestBool is template
func TakeWhileTestBool() string {
	return `
func TestTakeWhile<FTYPE>Ptr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []*<TYPE>{&vt, &vt, &vf}
	NewList := TakeWhile<FTYPE>Ptr(func(v *bool) bool { return *v == true }, []*<TYPE>{&vt, &vt, &vf, &vf, &vf})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("TakeWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*<TYPE>{&vt}
	NewList = TakeWhile<FTYPE>Ptr(func(v *bool) bool { return *v == true }, []*<TYPE>{&vt})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhile<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
	}

	if len(TakeWhile<FTYPE>Ptr(nil, []*<TYPE>{})) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}
`
}

// ReplaceActivityTakeWhilePtr replaces ...
func ReplaceActivityTakeWhilePtr(code string) string {
	s1 := `func TestTakeWhileStrPtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v40 string = "40"
	var v0 string

	expectedNewList := []*string{&v4, &v2, &v4}
	NewList, _ := TakeWhileStrPtrErr(isEvenStrPtrErr, []*string{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileStrPtrErr(isEvenStrPtrErr, []*string{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*string{&v40}
	partialIsEvenDivisibleBy := func(num *string) (bool, error) {
		if *num == "40" {
			return true, nil
		}
		return false, nil
	}
	NewList, _ = TakeWhileStrPtrErr(partialIsEvenDivisibleBy, []*string{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileStrPtrErr failed.")
	}

	r, _ = TakeWhileStrPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("TakeWhileStrPtr failed.")
	}
}`
	s2 := `func TestTakeWhileStrPtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v40 string = "40"
	var v0 string = "0"

	expectedNewList := []*string{&v4, &v2, &v4}
	NewList, _ := TakeWhileStrPtrErr(isEvenStrPtrErr, []*string{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileStrPtrErr(isEvenStrPtrErr, []*string{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*string{&v40}
	partialIsEvenDivisibleBy := func(num *string) (bool, error) {
		if *num == "40" {
			return true, nil
		}
		return false, nil
	}
	NewList, _ = TakeWhileStrPtrErr(partialIsEvenDivisibleBy, []*string{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileStrPtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileStrPtrErr failed.")
	}

	r, _ = TakeWhileStrPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("TakeWhileStrPtr failed.")
	}
}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestTakeWhileStrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v40 string = "40"
	var v0 string

	expectedNewList := []string{v4, v2, v4}
	NewList, _ := TakeWhileStrErr(isEvenStrErr, []string{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileStrErr(isEvenStrErr, []string{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []string{v40}
	partialIsEvenDivisibleBy := func(num string) (bool, error) { if num == "40" { return true, nil}; return false, nil }
	NewList, _ = TakeWhileStrErr(partialIsEvenDivisibleBy, []string{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileStrErr failed.")
	}

	r, _ = TakeWhileStrErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("TakeWhileStrErr failed.")
	}
}`

	s2 = `func TestTakeWhileStrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v40 string = "40"
	var v0 string = "0"

	expectedNewList := []string{v4, v2, v4}
	NewList, _ := TakeWhileStrErr(isEvenStrErr, []string{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhileStrErr(isEvenStrErr, []string{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []string{v40}
	partialIsEvenDivisibleBy := func(num string) (bool, error) { if num == "40" { return true, nil}; return false, nil }
	NewList, _ = TakeWhileStrErr(partialIsEvenDivisibleBy, []string{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileStrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhileStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhileStrErr failed.")
	}

	r, _ = TakeWhileStrErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("TakeWhileStrErr failed.")
	}
}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestTakeWhileStrPtr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v40 string = "40"

	expectedNewList := []*string{&v4, &v2, &v4}
	NewList := TakeWhileStrPtr(isEvenStrPtr, []*string{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*string{&v40}
	partialIsEvenDivisibleBy := func(num *string) bool { if *num == "40" { return true }; return false }
	NewList = TakeWhileStrPtr(partialIsEvenDivisibleBy, []*string{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileStrPtr(nil, nil)) > 0 {
		t.Errorf("TakeWhileStrPtr failed.")
	}

	if len(TakeWhileStrPtr(nil, []*string{})) > 0 {
		t.Errorf("TakeWhileStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}`

	s2 = `func TestTakeWhileStrPtr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v7 string = "7"
	var v40 string = "40"

	expectedNewList := []*string{&v4, &v2, &v4}
	NewList := TakeWhileStrPtr(isEvenStrPtr, []*string{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhileStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*string{&v40}
	partialIsEvenDivisibleBy := func(num *string) bool {
		if *num == "40" {
			return true
		}
		return false
	}
	NewList = TakeWhileStrPtr(partialIsEvenDivisibleBy, []*string{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhileStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileStrPtr(nil, nil)) > 0 {
		t.Errorf("TakeWhileStrPtr failed.")
	}

	if len(TakeWhileStrPtr(nil, []*string{})) > 0 {
		t.Errorf("TakeWhileStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}`
	code = strings.Replace(code, s1, s2, -1)
	return code
}

//**********************TakeWhile<TYPE>PtrErr******************

// TakeWhilePtrErrTest is template
func TakeWhilePtrErrTest() string {
	return `
func TestTakeWhile<FTYPE>PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v7 <TYPE> = 7
	var v40 <TYPE> = 40
	var v0 <TYPE>

	expectedNewList := []*<TYPE>{&v4, &v2, &v4}
	NewList, _ := TakeWhile<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{&v4, &v2, &v4, &v7, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("TakeWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhile<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{&v0, &v2, &v4, &v7, &v5})
	if err == nil {
		t.Errorf("TakeWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*<TYPE>{&v40}
	partialIsEvenDivisibleBy := func(num *<TYPE>) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhile<FTYPE>PtrErr(partialIsEvenDivisibleBy, []*<TYPE>{&v40})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhile<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhile<FTYPE>PtrErr failed.")
	}

	r, _ = TakeWhile<FTYPE>PtrErr(nil, []*<TYPE>{})
	if len(r) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
	}
}
`
}

// TakeWhilePtrErrTestBool is template
func TakeWhilePtrErrTestBool() string {
	return `
func TestTakeWhile<FTYPE>PtrErr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []*<TYPE>{&vt, &vt, &vf}
	NewList, _ := TakeWhile<FTYPE>PtrErr(func(v *bool) (bool, error) { return *v == true, nil }, []*<TYPE>{&vt, &vt, &vf, &vf, &vf})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("TakeWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []*<TYPE>{&vt}
	NewList, _ = TakeWhile<FTYPE>PtrErr(func(v *bool) (bool, error) { return *v == true, nil }, []*<TYPE>{&vt})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("TakeWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhile<FTYPE>PtrErr(func(v *bool) (bool, error) {
		if *v == false {
			return false, errors.New("false is invalid for this test")
		}
		return *v == true, nil
	}, []*<TYPE>{&vf})

	if err == nil {
		t.Errorf("TakeWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhile<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhile<FTYPE>PtrErr failed.")
	}

	r, _ = TakeWhile<FTYPE>PtrErr(nil, []*<TYPE>{})
	if len(r) > 0 {
		t.Errorf("TakeWhile<FTYPE>Ptr failed.")
	}
}
`
}

// ReplaceActivityTakeWhilePtrErr replaces ...
func ReplaceActivityTakeWhilePtrErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestTakeWhileIntPtrErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestTakeWhileIntPtrErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEvenDivisibleBy := func(num *string) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileStrPtrErr(partialIsEvenDivisibleBy, []*string{&v40})`

	s2 = `partialIsEvenDivisibleBy := func(num *string) (bool, error) {
		if *num == "40" {
			return true, nil
		}
		return false, nil
	}
	NewList, _ = TakeWhileStrPtrErr(partialIsEvenDivisibleBy, []*string{&v40})`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEvenDivisibleBy := func(num *float64) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileFloat64PtrErr(partialIsEvenDivisibleBy, []*float64{&v40})`

	s2 = `partialIsEvenDivisibleBy := func(num *float64) (bool, error) { return int(*num)%10 == 0, nil }
	NewList, _ = TakeWhileFloat64PtrErr(partialIsEvenDivisibleBy, []*float64{&v40})`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEvenDivisibleBy := func(num *float32) (bool, error) { return *num%10 == 0, nil }
	NewList, _ = TakeWhileFloat32PtrErr(partialIsEvenDivisibleBy, []*float32{&v40})`

	s2 = `partialIsEvenDivisibleBy := func(num *float32) (bool, error) { return int(*num)%10 == 0, nil }
	NewList, _ = TakeWhileFloat32PtrErr(partialIsEvenDivisibleBy, []*float32{&v40})`

	code = strings.Replace(code, s1, s2, -1)
	return code
}

//**********************TakeWhile<TYPE>Err******************

// TakeWhileErrTest is template
func TakeWhileErrTest() string {
	return `
func TestTakeWhile<FTYPE>Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v7 <TYPE> = 7
	var v40 <TYPE> = 40
	var v0 <TYPE>

	expectedNewList := []<TYPE>{v4, v2, v4}
	NewList, _ := TakeWhile<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{v4, v2, v4, v7, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhile<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{v0, v2, v4, v7, v5})
	if err == nil {
		t.Errorf("TakeWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []<TYPE>{v40}
	partialIsEvenDivisibleBy := func(num <TYPE>) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhile<FTYPE>Err(partialIsEvenDivisibleBy, []<TYPE>{v40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhile<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhile<FTYPE>Err failed.")
	}

	r, _ = TakeWhile<FTYPE>Err(nil, []<TYPE>{})
	if len(r) > 0 {
		t.Errorf("TakeWhile<FTYPE>Err failed.")
	}
}
`
}

// TakeWhileErrTestBool is template
func TakeWhileErrTestBool() string {
	return `
func TestTakeWhile<FTYPE>Err(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []<TYPE>{vt, vt, vf}
	NewList, _ := TakeWhile<FTYPE>Err(func(v bool) (bool, error) {return v == true, nil}, []<TYPE>{vt, vt, vf, vf, vf})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("TakeWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []<TYPE>{vt}
	NewList, _ = TakeWhile<FTYPE>Err(func(v bool) (bool, error) {return v == true, nil}, []<TYPE>{vt})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := TakeWhile<FTYPE>Err(func(v bool) (bool, error) { if v == false { return false, errors.New("false is invalid for this test") }; return v == true, nil}, []<TYPE>{vf})
	
	if err == nil {
		t.Errorf("TakeWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := TakeWhile<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("TakeWhile<FTYPE>Err failed.")
	}

	r, _ = TakeWhile<FTYPE>Err(nil, []<TYPE>{})
	if len(r) > 0 {
		t.Errorf("TakeWhile<FTYPE>Err failed.")
	}
}
`
}

// ReplaceActivityTakeWhileErr replaces ...
func ReplaceActivityTakeWhileErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestTakeWhileIntErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestTakeWhileIntErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEvenDivisibleBy := func(num string) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileStrErr(partialIsEvenDivisibleBy, []string{v40})`

	s2 = `partialIsEvenDivisibleBy := func(num string) (bool, error) { if num == "40" { return true, nil}; return false, nil }
	NewList, _ = TakeWhileStrErr(partialIsEvenDivisibleBy, []string{v40})`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEvenDivisibleBy := func(num float64) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileFloat64Err(partialIsEvenDivisibleBy, []float64{v40})`

	s2 = `partialIsEvenDivisibleBy := func(num float64) (bool, error) { return int(num)%10 == 0, nil }
	NewList, _ = TakeWhileFloat64Err(partialIsEvenDivisibleBy, []float64{v40})`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEvenDivisibleBy := func(num float32) (bool, error) { return num%10 == 0, nil }
	NewList, _ = TakeWhileFloat32Err(partialIsEvenDivisibleBy, []float32{v40})`

	s2 = `partialIsEvenDivisibleBy := func(num float32) (bool, error) { return int(num)%10 == 0, nil }
	NewList, _ = TakeWhileFloat32Err(partialIsEvenDivisibleBy, []float32{v40})`

	code = strings.Replace(code, s1, s2, -1)
	return code
}
