package basic

import "strings"

// SomePtrTest is template
func SomePtrTest() string {
	return `
func TestSome<FTYPE>Ptr(t *testing.T) {
	// Test : value exist in the list

	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10

	list1 := []*<TYPE>{&v8, &v2, &v10, &v4}
	if !Some<FTYPE>Ptr(isEven<FTYPE>Ptr, list1) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	list2 := []*<TYPE>{&v1, &v3, &v5, &v7}
	if Some<FTYPE>Ptr(isEven<FTYPE>Ptr, list2) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Some<FTYPE>Ptr(nil, nil) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Some<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{}) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
		t.Errorf(reflect.String.String())
	}
}
`
}

// SomePtrTestBool is template
func SomePtrTestBool() string {
	return `
func TestSome<FTYPE>Ptr(t *testing.T) {
	// Test : value exist in the list

	var vt <TYPE> = true
	var vf <TYPE> = false

	list1 := []*<TYPE>{&vt, &vf}
	if !Some<FTYPE>Ptr(func(v *bool) bool { return *v == true }, list1) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	if Some<FTYPE>Ptr(nil, nil) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	if Some<FTYPE>Ptr(func(v *bool) bool { return *v == true }, []*<TYPE>{}) {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}
`
}

//************************SomePtrErr Function*******************

// SomePtrErrTest is template
func SomePtrErrTest() string {
	return `
func TestSome<FTYPE>PtrErr(t *testing.T) {
	// Test : value exist in the list

	var v0 <TYPE> = 0
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10

	list1 := []*<TYPE>{&v8, &v2, &v10, &v4}
	r, _ := Some<FTYPE>PtrErr(isEven<FTYPE>PtrErr, list1)
	if !r {
		t.Errorf("Some<FTYPE>PtrErr failed. Expected=true, actual=false")
	}

	list3 := []*<TYPE>{&v0, &v4}
	_, err := Some<FTYPE>PtrErr(isEven<FTYPE>PtrErr, list3)
	if err == nil {
		t.Errorf("Some<FTYPE>PtrErr failed. Expected=true, actual=false")
	}

	list2 := []*<TYPE>{&v1, &v3, &v5, &v7}
	r, _ = Some<FTYPE>PtrErr(isEven<FTYPE>PtrErr, list2)
	if r {
		t.Errorf("Some<FTYPE>PtrErr failed. Expected=false, actual=true")
	}

	r, _ = Some<FTYPE>PtrErr(nil, nil)
	if r {
		t.Errorf("Some<FTYPE>PtrErr failed. Expected=false, actual=true")
	}

	r, _ = Some<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{})
	if r {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}
`
}

// SomePtrErrTestBool is template
func SomePtrErrTestBool() string {
	return `
func TestSome<FTYPE>PtrErr(t *testing.T) {
	// Test : value exist in the list

	var vt <TYPE> = true
	var vf <TYPE> = false

	list1 := []*<TYPE>{&vt, &vf}
	r, _ := Some<FTYPE>PtrErr(func(v *bool) (bool, error) { return *v == true, nil }, list1)
	if !r {
		t.Errorf("Some<FTYPE>PtrErr failed. Expected=true, actual=false")
	}

	r, _ = Some<FTYPE>PtrErr(nil, nil)

	if r {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}

	r, _ = Some<FTYPE>PtrErr(func(v *bool) (bool, error) { return *v == true, nil }, []*<TYPE>{})
	if r {
		t.Errorf("Some<FTYPE>PtrErr failed. Expected=false, actual=true")
	}

	_, err := Some<FTYPE>PtrErr(func(v *bool) (bool, error) {
		if *v == false {
			return false, errors.New("false is invalid in this test")
		}
		return *v == true, nil
	}, []*<TYPE>{&vf})

	if err == nil {
		t.Errorf("Some<FTYPE>PtrErr failed. Expected=false, actual=true")
	}
}
`
}

// ReplaceActivitySomePtrErr replaces ...
func ReplaceActivitySomePtrErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSomeIntPtrErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestSomeIntPtrErr(t *testing.T) {`

	return strings.Replace(code, s1, s2, -1)
}

//************************SomeErr Function*******************

// SomeErrTest is template
func SomeErrTest() string {
	return `
func TestSome<FTYPE>Err(t *testing.T) {
	// Test : value exist in the list

	var v0 <TYPE> = 0
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v7 <TYPE> = 7
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10

	list1 := []<TYPE>{v8, v2, v10, v4}
	r, _ := Some<FTYPE>Err(isEven<FTYPE>Err, list1)
	if !r {
		t.Errorf("Some<FTYPE>Err failed. Expected=true, actual=false")
	}

	list3 := []<TYPE>{v0, v4}
	_, err := Some<FTYPE>Err(isEven<FTYPE>Err, list3)
	if err == nil {
		t.Errorf("Some<FTYPE>Err failed. Expected=true, actual=false")
	}

	list2 := []<TYPE>{v1, v3, v5, v7}
	r, _ = Some<FTYPE>Err(isEven<FTYPE>Err, list2)
	if r {
		t.Errorf("Some<FTYPE>Err failed. Expected=false, actual=true")
	}

	r, _ = Some<FTYPE>Err(nil, nil)
	if r {
		t.Errorf("Some<FTYPE>Err failed. Expected=false, actual=true")
	}

	r, _ = Some<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{})
	if r {
		t.Errorf("Some<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}
`
}

// SomeErrTestBool is template
func SomeErrTestBool() string {
	return `
func TestSome<FTYPE>Err(t *testing.T) {
	// Test : value exist in the list

	var vt <TYPE> = true
	var vf <TYPE> = false

	list1 := []<TYPE>{vt, vf}
	r, _ := Some<FTYPE>Err(func(v bool) (bool, error) { return v == true, nil }, list1)
	if !r {
		t.Errorf("Some<FTYPE> Err failed. Expected=true, actual=false")
	}

	r, _ = Some<FTYPE>Err(nil, nil)
	
	if r{
		t.Errorf("Some<FTYPE>Err failed. Expected=false, actual=true")
	}

	r, _ = Some<FTYPE>Err(func(v bool) (bool, error) { return v == true, nil }, []<TYPE>{})
	if  r {
		t.Errorf("Some<FTYPE>Err failed. Expected=false, actual=true")
	}

	_, err := Some<FTYPE>Err(func(v bool) (bool, error) { if v == false { return false, errors.New("false is invalid in this test") }; return v == true, nil }, []<TYPE>{vf})
	if  err == nil {
		t.Errorf("Some<FTYPE>Err failed. Expected=false, actual=true")
	}
}
`
}

// ReplaceActivitySomeErr replaces ...
func ReplaceActivitySomeErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSomeIntErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestSomeIntErr(t *testing.T) {`

	return strings.Replace(code, s1, s2, -1)
}
