package basic

import "strings"

// Drop returns a new list after dropping the given item
func DropWhilePtrTest() string {
	return `
func TestDropWhile<FTYPE>Ptr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	expectedNewList := []*<TYPE>{&v3, &v4, &v5}
	NewList := DropWhile<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhile<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
	}

	if len(DropWhile<FTYPE>Ptr(nil, []*<TYPE>{})) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}

	NewList = DropWhile<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed")
	}
}

`
}

// DistinctBoolPtr removes duplicates.
func DropWhilePtrBoolTest() string {
	return `
func TestDropWhile<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []*<TYPE>{&vf, &vt}
	NewList := DropWhile<FTYPE>Ptr(isTrue<FTYPE>Ptr, []*<TYPE>{&vt, &vf, &vt})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("DropWhile<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}

`
}

//**********DropWhilePtrErrTest************************
// Drop returns a new list after dropping the given item
func DropWhilePtrErrTest() string {
	return `
func TestDropWhile<FTYPE>PtrErr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v0 <TYPE> = 0

	expectedNewList := []*<TYPE>{&v3, &v4, &v5}
	NewList, _ := DropWhile<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhile<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{&v4, &v2, &v0, &v4, &v5})
	if err == nil {
		t.Errorf("DropWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhile<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
	}

	r, _ = DropWhile<FTYPE>PtrErr(nil, []*<TYPE>{})
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
	}

	NewList, _ = DropWhile<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{&v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhile<FTYPE>PtrErr failed")
	}
}

`
}

// DropWhilePtrErrBoolTest
func DropWhilePtrErrBoolTest() string {
	return `
func TestDropWhile<FTYPE>PtrErr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []*<TYPE>{&vf, &vt}
	NewList, _ := DropWhile<FTYPE>PtrErr(isTrue<FTYPE>PtrErr, []*<TYPE>{&vt, &vf, &vt})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("DropWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhile<FTYPE>PtrErr(isTrue<FTYPE>PtrErr2, []*<TYPE>{&vt, &vf, &vt})
	if err == nil {
		t.Errorf("DropWhile<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhile<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
	}

	r, _ = DropWhile<FTYPE>PtrErr(nil, []*<TYPE>{})
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
	}

	r, _ = DropWhile<FTYPE>PtrErr(isTrue<FTYPE>PtrErr2, []*<TYPE>{})
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Ptr failed.")
	}
}

func isTrueBoolPtrErr(v *bool) (bool, error) {
	return *v == true, nil
}
func isTrueBoolPtrErr2(v *bool) (bool, error) {
	if *v == false {
		return false, errors.New("false is invalid for this test")
	}
	return *v == true, nil
}
`
}

func ReplaceActivityDropWhilePtrErr(code string) string {
	s1 := `import (
    _ "errors"
	"reflect"
	"testing"
)

func TestDropWhileIntPtrErr(t *testing.T) {`
	s2 := `import (
    "errors"
	"testing"
)

func TestDropWhileIntPtrErr(t *testing.T) {`

	return strings.Replace(code, s1, s2, -1)
}

//**********DropWhileErrTest************************
// Drop returns a new list after dropping the given item
func DropWhileErrTest() string {
	return `
func TestDropWhile<FTYPE>Err(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v0 <TYPE> = 0

	expectedNewList := []<TYPE>{v3, v4, v5}
	NewList, _ := DropWhile<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{v4, v2, v3, v4, v5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhile<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{v4, v2, v0, v4, v5})
	if err == nil {
		t.Errorf("DropWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhile<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Err failed.")
	}

	r, _ = DropWhile<FTYPE>Err(nil, []<TYPE>{})
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Err failed.")
	}

	NewList, _ = DropWhile<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{v4})
	if len(NewList) != 0 {
		t.Errorf("DropWhile<FTYPE>Err failed")
	}
}

`
}

// DropWhileErrBoolTest
func DropWhileErrBoolTest() string {
	return `
func TestDropWhile<FTYPE>Err(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []<TYPE>{vf, vt}
	NewList, _ := DropWhile<FTYPE>Err(isTrue<FTYPE>Err, []<TYPE>{vt, vf, vt})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("DropWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	_, err := DropWhile<FTYPE>Err(isTrue<FTYPE>Err2, []<TYPE>{vt, vf, vt})
	if err == nil {
		t.Errorf("DropWhile<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := DropWhile<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Err failed.")
	}

	r, _ = DropWhile<FTYPE>Err(nil, []<TYPE>{})
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Err failed.")
	}

	r, _ = DropWhile<FTYPE>Err(isTrue<FTYPE>Err2, []<TYPE>{})
	if len(r) > 0 {
		t.Errorf("DropWhile<FTYPE>Err failed.")
	}
}

func isTrueBoolErr(v bool) (bool, error) {
	return v == true, nil
}
func isTrueBoolErr2(v bool) (bool, error) {
	if v == false {
		return false, errors.New("false is invalid for this test")
	}
	return v == true, nil
}
`
}

func ReplaceActivityDropWhileErr(code string) string {
	s1 := `import (
    _ "errors"
	"reflect"
	"testing"
)

func TestDropWhileIntErr(t *testing.T) {`
	s2 := `import (
    "errors"
	"testing"
)

func TestDropWhileIntErr(t *testing.T) {`

	return strings.Replace(code, s1, s2, -1)
}
