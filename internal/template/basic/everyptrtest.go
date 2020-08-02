package basic

import "strings"

// EveryPtrTest applies the function(1st argument) on each item of the list and returns new list
func EveryPtrTest() string {
	return `
func TestEvery<FTYPE>Ptr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10
	list1 := []*<TYPE>{&v8, &v2, &v10, &v4}
	if !Every<FTYPE>Ptr(isEven<FTYPE>Ptr, list1) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	list2 := []*<TYPE>{&v8, &v2, &v10, &v5, &v4}
	if Every<FTYPE>Ptr(isEven<FTYPE>Ptr, list2) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Every<FTYPE>Ptr(isEven<FTYPE>Ptr, nil) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Every<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{}) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Every<FTYPE>Ptr(nil, []*<TYPE>{}) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
`
}

// EveryPtrBoolTest applies the function(1st argument) on each item of the list and returns new list
func EveryPtrBoolTest() string {
	return `
func TestEvery<FTYPE>Ptr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	// Test : every value in the list is either true or false
	list1 := []*<TYPE>{&vt, &vt, &vt, &vt}
	if !Every<FTYPE>Ptr(TruePtr, list1) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	list1 = []*<TYPE>{&vt, &vt, &vt, &vf}
	if Every<FTYPE>Ptr(TruePtr, list1) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	list1 = []*bool{}
	if Every<FTYPE>Ptr(TruePtr, list1) {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	if EveryBoolPtr(TruePtr, nil) {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	if EveryBoolPtr(nil, []*bool{}) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}
	if EveryBoolPtr(TruePtr, []*bool{}) {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}

func TruePtr(val *bool) bool {
	return *val == true
}
`
}

// ReplaceActivityEveryPtrTest replaces ...
func ReplaceActivityEveryPtrTest(code string) string {
	s1 := `func TestEveryStrPtr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v8 string = "8"
	var v10 string = "10"
	list1 := []*string{&v8, &v2, &v10, &v4}
	if !EveryStrPtr(isEvenStrPtr, list1) {
		t.Errorf("EveryStrPtr failed. Expected=true, actual=false")
	}

	list2 := []*string{&v8, &v2, &v10, &v5, &v4}
	if EveryStrPtr(isEvenStrPtr, list2) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}

	if EveryStrPtr(isEvenStrPtr, nil) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}

	if EveryStrPtr(isEvenStrPtr, []*string{}) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}

	if EveryStrPtr(nil, []*string{}) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}`
	s2 := `func TestEveryStrPtr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	list1 := []*string{&v2, &v4}
	if !EveryStrPtr(isEvenStrPtr, list1) {
		t.Errorf("EveryStrPtr failed. Expected=true, actual=false")
	}

	list1 = []*string{&v2, &v4, &v3}
	if EveryStrPtr(isEvenStrPtr, list1) {
		t.Errorf("EveryStrPtr failed. Expected=true, actual=false")
	}

	if EveryStrPtr(isEvenStrPtr, nil) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}

	if EveryStrPtr(isEvenStrPtr, []*string{}) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}

	if EveryStrPtr(nil, []*string{}) {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}

//*************EveryPtrErr******************

// EveryPtrErrTest applies the function(1st argument) on each item of the list and returns new list
func EveryPtrErrTest() string {
	return `
func TestEvery<FTYPE>PtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10
	var v0 <TYPE>

	list1 := []*<TYPE>{&v2, &v4}

	r, _ := Every<FTYPE>PtrErr(isEven<FTYPE>PtrErr, list1)
	if !r {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*<TYPE>{&v0, &v4}
	_, err := Every<FTYPE>PtrErr(isEven<FTYPE>PtrErr, list1)
	if err == nil {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*<TYPE>{&v8, &v2, &v10, &v5, &v4}
	r, _ = Every<FTYPE>PtrErr(isEven<FTYPE>PtrErr, list2)
	if r {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=false, actual=true")
	}

	r, _ = Every<FTYPE>PtrErr(isEven<FTYPE>PtrErr, nil)
	if r {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=false, actual=true")
	}

	r, _ = Every<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{})
	if r {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=false, actual=true")
	}

	r, _ = Every<FTYPE>PtrErr(nil, []*<TYPE>{})
	if r {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}
`
}

// EveryPtrErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func EveryPtrErrBoolTest() string {
	return `
func TestEvery<FTYPE>PtrErr(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	// Test : every value in the list is either true or false
	list1 := []*<TYPE>{&vt, &vt, &vt, &vt}
	r, _ := Every<FTYPE>PtrErr(TruePtrErr, list1)
	if !r {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*<TYPE>{&vt, &vt, &vt}
	r, _ = Every<FTYPE>PtrErr(TruePtrErr, list1)
	if !r {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*<TYPE>{&vt, &vt, &vt, &vf}
	_, err := Every<FTYPE>PtrErr(TruePtrErr, list1)
	if err == nil {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*<TYPE>{&vt, &vt, &vf}
	r, _ = Every<FTYPE>PtrErr(TruePtrErr2, list1)
	if r {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=true, actual=false")
	}

	list1 = []*bool{}
	r, _ = Every<FTYPE>PtrErr(TruePtrErr, list1)
	if r {
		t.Errorf("EveryBool failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolPtrErr(TruePtrErr, nil)
	if r {
		t.Errorf("EveryBoolPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolPtrErr(nil, []*bool{})
	if r {
		t.Errorf("Every<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolPtrErr(TruePtrErr, []*bool{})
	if r {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=false, actual=true")
	}
}

func TruePtrErr(val *bool) (bool, error) {
	if *val == false {
		return false, errors.New("false is invalid in this test")
	}
	return *val == true, nil
}

func TruePtrErr2(val *bool) (bool, error) {
	return *val == true, nil
}
`
}

// ReplaceActivityEveryPtrErr replaces ...
func ReplaceActivityEveryPtrErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestEveryIntPtrErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestEveryIntPtrErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestEveryStrPtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v8 string = "8"
	var v10 string = "10"
	var v0 string

	list1 := []*string{&v2, &v4}

	r, _ := EveryStrPtrErr(isEvenStrPtrErr, list1)
	if !r {
		t.Errorf("EveryStrPtrErr failed. Expected=true, actual=false")
	}

	list1 = []*string{&v0, &v4}
	_, err := EveryStrPtrErr(isEvenStrPtrErr, list1)
	if err == nil {
		t.Errorf("EveryStrPtrErr failed. Expected=true, actual=false")
	}

	list2 := []*string{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryStrPtrErr(isEvenStrPtrErr, list2)
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrPtrErr(isEvenStrPtrErr, nil)
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrPtrErr(isEvenStrPtrErr, []*string{})
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrPtrErr(nil, []*string{})
	if r {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}
}`
	s2 = `func TestEveryStrPtrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v8 string = "8"
	var v10 string = "10"
	var v0 string = "0"

	list1 := []*string{&v2, &v4}

	r, _ := EveryStrPtrErr(isEvenStrPtrErr, list1)
	if !r {
		t.Errorf("EveryStrPtrErr failed. Expected=true, actual=false")
	}

	list1 = []*string{&v0, &v4}
	_, err := EveryStrPtrErr(isEvenStrPtrErr, list1)
	if err == nil {
		t.Errorf("EveryStrPtrErr failed. Expected=true, actual=false")
	}

	list2 := []*string{&v8, &v2, &v10, &v5, &v4}
	r, _ = EveryStrPtrErr(isEvenStrPtrErr, list2)
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrPtrErr(isEvenStrPtrErr, nil)
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrPtrErr(isEvenStrPtrErr, []*string{})
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrPtrErr(nil, []*string{})
	if r {
		t.Errorf("EveryStrPtr failed. Expected=false, actual=true")
	}
}`

	code = strings.Replace(code, s1, s2, -1)

	return code
}

//*************EveryErr******************

// EveryErrTest applies the function(1st argument) on each item of the list and returns new list
func EveryErrTest() string {
	return `
func TestEvery<FTYPE>Err(t *testing.T) {
	// Test : every value in the list is even number
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10
	var v0 <TYPE>

	list1 := []<TYPE>{v2, v4}

	r, _ := Every<FTYPE>Err(isEven<FTYPE>Err, list1)
	if !r {
		t.Errorf("Every<FTYPE>Err failed. Expected=true, actual=false")
	}

	list1 = []<TYPE>{v0, v4}
	_, err := Every<FTYPE>Err(isEven<FTYPE>Err, list1)
	if err == nil {
		t.Errorf("Every<FTYPE>Err failed. Expected=true, actual=false")
	}

	list2 := []<TYPE>{v8, v2, v10, v5, v4}
	r, _ = Every<FTYPE>Err(isEven<FTYPE>Err, list2)
	if r {
		t.Errorf("Every<FTYPE>PtrErr failed. Expected=false, actual=true")
	}

	r, _ = Every<FTYPE>Err(isEven<FTYPE>Err, nil)
	if r {
		t.Errorf("Every<FTYPE>Err failed. Expected=false, actual=true")
	}

	r, _ = Every<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{})
	if r {
		t.Errorf("Every<FTYPE>Err failed. Expected=false, actual=true")
	}

	r, _ = Every<FTYPE>Err(nil, []<TYPE>{})
	if r {
		t.Errorf("Every<FTYPE>Err failed. Expected=false, actual=true")
	}
}
`
}

// EveryErrBoolTest applies the function(1st argument) on each item of the list and returns new list
func EveryErrBoolTest() string {
	return `
func TestEvery<FTYPE>Err(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	// Test : every value in the list is either true or false
	list1 := []<TYPE>{vt, vt, vt, vt}
	r, _ := Every<FTYPE>Err(TrueErr, list1)
	if !r {
		t.Errorf("Every<FTYPE>Err failed. Expected=true, actual=false")
	}

	list1 = []<TYPE>{vt, vt, vt}
	r, _ = Every<FTYPE>Err(TrueErr, list1)
	if !r {
		t.Errorf("Every<FTYPE>Err failed. Expected=true, actual=false")
	}

	list1 = []<TYPE>{vt, vt, vt, vf}
	_, err := Every<FTYPE>Err(TrueErr, list1)
	if err == nil {
		t.Errorf("Every<FTYPE>Err failed. Expected=true, actual=false")
	}

	list1 = []<TYPE>{vt, vt, vf}
	r, _ = Every<FTYPE>Err(TrueErr2, list1)
	if r {
		t.Errorf("Every<FTYPE>Err failed. Expected=true, actual=false")
	}

	list1 = []bool{}
	r, _ = Every<FTYPE>Err(TrueErr, list1)
	if r {
		t.Errorf("EveryBoolErr failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolErr(TrueErr, nil)
	if r {
		t.Errorf("EveryBoolErr failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolErr(nil, []bool{})
	if r {
		t.Errorf("Every<FTYPE>Err failed. Expected=false, actual=true")
	}

	r, _ = EveryBoolErr(TrueErr, []bool{})
	if r {
		t.Errorf("Every<FTYPE>Err failed. Expected=false, actual=true")
	}
}

func TrueErr(val bool) (bool, error) {
	if val == false {
		return false, errors.New("false is invalid in this test")
	}
	return val == true, nil
}

func TrueErr2(val bool) (bool, error) {
	return val == true, nil
}
`
}

// ReplaceActivityEveryErr replaces ...
func ReplaceActivityEveryErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestEveryIntErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestEveryIntPErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestEveryStrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v8 string = "8"
	var v10 string = "10"
	var v0 string

	list1 := []string{v2, v4}

	r, _ := EveryStrErr(isEvenStrErr, list1)
	if !r {
		t.Errorf("EveryStrErr failed. Expected=true, actual=false")
	}

	list1 = []string{v0, v4}
	_, err := EveryStrErr(isEvenStrErr, list1)
	if err == nil {
		t.Errorf("EveryStrErr failed. Expected=true, actual=false")
	}

	list2 := []string{v8, v2, v10, v5, v4}
	r, _ = EveryStrErr(isEvenStrErr, list2)
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrErr(isEvenStrErr, nil)
	if r {
		t.Errorf("EveryStrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrErr(isEvenStrErr, []string{})
	if r {
		t.Errorf("EveryStrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrErr(nil, []string{})
	if r {
		t.Errorf("EveryStrErr failed. Expected=false, actual=true")
	}
}`
	s2 = `func TestEveryStrErr(t *testing.T) {
	// Test : every value in the list is even number
	var v2 string = "2"
	var v4 string = "4"
	var v5 string = "5"
	var v8 string = "8"
	var v10 string = "10"
	var v0 string = "0"

	list1 := []string{v2, v4}

	r, _ := EveryStrErr(isEvenStrErr, list1)
	if !r {
		t.Errorf("EveryStrErr failed. Expected=true, actual=false")
	}

	list1 = []string{v0, v4}
	_, err := EveryStrErr(isEvenStrErr, list1)
	if err == nil {
		t.Errorf("EveryStrErr failed. Expected=true, actual=false")
	}

	list2 := []string{v8, v2, v10, v5, v4}
	r, _ = EveryStrErr(isEvenStrErr, list2)
	if r {
		t.Errorf("EveryStrPtrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrErr(isEvenStrErr, nil)
	if r {
		t.Errorf("EveryStrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrErr(isEvenStrErr, []string{})
	if r {
		t.Errorf("EveryStrErr failed. Expected=false, actual=true")
	}

	r, _ = EveryStrErr(nil, []string{})
	if r {
		t.Errorf("EveryStrErr failed. Expected=false, actual=true")
	}
}`

	code = strings.Replace(code, s1, s2, -1)

	return code
}
