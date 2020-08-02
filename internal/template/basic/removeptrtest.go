package basic

import "strings"

// RemovePtrTest is template
func RemovePtrTest() string {
	return `
func TestRemove<FTYPE>Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v20 <TYPE> = 20
	var v40 <TYPE> = 40

	expectedNewList := []*<TYPE>{&v1, &v3}
	NewList := Remove<FTYPE>Ptr(isEven<FTYPE>Ptr, []*<TYPE>{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("Remove<FTYPE> failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*<TYPE>{&v1, &v3}
	partialIsEven := func(num *<TYPE>) bool {
		return *num%10 == 0
	}
	NewList = Remove<FTYPE>Ptr(partialIsEven, []*<TYPE>{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("Remove<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(Remove<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Remove<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}
`
}

// RemovePtrTestBool is template
func RemovePtrTestBool() string {
	return `
func TestRemove<FTYPE>Ptr(t *testing.T) {
	// Test : even number in the list
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []*<TYPE>{&vt}
	NewList := Remove<FTYPE>Ptr(func(v *bool) bool { return *v == false }, []*<TYPE>{&vt, &vf, &vf})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("Remove<FTYPE> failed. Expected New list=%v, actual list=%v", *expectedNewList[0], *NewList[0])
	}

	if len(Remove<FTYPE>Ptr(nil, nil)) > 0 {
		t.Errorf("Remove<FTYPE>Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}
`
}

// ReplaceActivityRemoveStrPtr replaces ...
func ReplaceActivityRemoveStrPtr(code string) string {
	s1 := `unc TestRemoveStrPtr(t *testing.T) {
	// Test : even number in the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v20 string = "20"
	var v40 string = "40"

	expectedNewList := []*string{&v1, &v3}
	NewList := RemoveStrPtr(isEvenStrPtr, []*string{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*string{&v1, &v3}
	partialIsEven := func(num *string) bool {
		return *num%10 == 0
	}
	NewList = RemoveStrPtr(partialIsEven, []*string{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveStrPtr(nil, nil)) > 0 {
		t.Errorf("RemoveStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}`
	s2 := `unc TestRemoveStrPtr(t *testing.T) {
	// Test : even number in the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v20 string = "20"
	var v40 string = "40"

	expectedNewList := []*string{&v1, &v3}
	NewList := RemoveStrPtr(isEvenStrPtr, []*string{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*string{&v1, &v3}
	partialIsEven := func(num *string) bool {
		if *num == "10" || *num == "20" {
			return true
		}
		return false
	}
	NewList = RemoveStrPtr(partialIsEven, []*string{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveStrPtr(nil, nil)) > 0 {
		t.Errorf("RemoveStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestRemoveFloat32Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v20 float32 = 20
	var v40 float32 = 40

	expectedNewList := []*float32{&v1, &v3}
	NewList := RemoveFloat32Ptr(isEvenFloat32Ptr, []*float32{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*float32{&v1, &v3}
	partialIsEven := func(num *float32) bool {
		return *num%10 == 0
	}
	NewList = RemoveFloat32Ptr(partialIsEven, []*float32{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveFloat32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}`
	s2 = `func TestRemoveFloat32Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v20 float32 = 20
	var v40 float32 = 40

	expectedNewList := []*float32{&v1, &v3}
	NewList := RemoveFloat32Ptr(isEvenFloat32Ptr, []*float32{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*float32{&v1, &v3}
	partialIsEven := func(num *float32) bool {
		return int(*num)%10 == 0
	}
	NewList = RemoveFloat32Ptr(partialIsEven, []*float32{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat32Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveFloat32Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestRemoveFloat64Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v20 float64 = 20
	var v40 float64 = 40

	expectedNewList := []*float64{&v1, &v3}
	NewList := RemoveFloat64Ptr(isEvenFloat64Ptr, []*float64{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*float64{&v1, &v3}
	partialIsEven := func(num *float64) bool {
		return *num%10 == 0
	}
	NewList = RemoveFloat64Ptr(partialIsEven, []*float64{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveFloat64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}`

	s2 = `func TestRemoveFloat64Ptr(t *testing.T) {
	// Test : even number in the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v20 float64 = 20
	var v40 float64 = 40

	expectedNewList := []*float64{&v1, &v3}
	NewList := RemoveFloat64Ptr(isEvenFloat64Ptr, []*float64{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*float64{&v1, &v3}
	partialIsEven := func(num *float64) bool {
		return int(*num)%10 == 0
	}
	NewList = RemoveFloat64Ptr(partialIsEven, []*float64{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("RemoveFloat64Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("RemoveFloat64Ptr failed.")
		t.Errorf(reflect.String.String())
	}
}`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEvenDivisibleBy := func(num *float32) bool { return *num%10 == 0 }`
	s2 = `partialIsEvenDivisibleBy := func(num *float32) bool { return int(*num)%10 == 0 }`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEvenDivisibleBy := func(num *float64) bool { return *num%10 == 0 }`
	s2 = `partialIsEvenDivisibleBy := func(num *float64) bool { return int(*num)%10 == 0 }`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEvenDivisibleBy := func(num *string) bool { return *num%10 == 0 }`
	s2 = `partialIsEvenDivisibleBy := func(num *string) bool { if *num == "40" { return true }; return false }`
	code = strings.Replace(code, s1, s2, -1)
	return code
}

//**********RemovePtrErr***********************

// RemovePtrErrTest is template
func RemovePtrErrTest() string {
	return `
func TestRemove<FTYPE>PtrErr(t *testing.T) {
	// Test : even number in the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v20 <TYPE> = 20
	var v40 <TYPE> = 40

	expectedNewList := []*<TYPE>{&v1, &v3}
	NewList, _ := Remove<FTYPE>PtrErr(isEven<FTYPE>PtrErr, []*<TYPE>{&v1, &v2, &v3, &v4})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("Remove<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []*<TYPE>{&v1, &v3}
	partialIsEven := func(num *<TYPE>) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = Remove<FTYPE>PtrErr(partialIsEven, []*<TYPE>{&v20, &v1, &v3, &v40})

	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] {
		t.Errorf("Remove<FTYPE>Ptr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := Remove<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Remove<FTYPE>Ptr failed.")
	}

	_, err := Remove<FTYPE>PtrErr(func(v *<TYPE>) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*<TYPE>{&v20, &v1, &v3, &v40})
	if err == nil {
		t.Errorf("Remove<FTYPE>PtrErr failed.")
	}
}
`
}

// RemovePtrErrTestBool is template
func RemovePtrErrTestBool() string {
	return `
func TestRemove<FTYPE>PtrErr(t *testing.T) {
	// Test : even number in the list
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []*<TYPE>{&vt}
	NewList, _ := Remove<FTYPE>PtrErr(func(v *bool) (bool, error) { return *v == false, nil }, []*<TYPE>{&vt, &vf, &vf})

	if *NewList[0] != *expectedNewList[0] {
		t.Errorf("Remove<FTYPE>PtrErr failed. Expected New list=%v, actual list=%v", *expectedNewList[0], *NewList[0])
	}

	_, err := Remove<FTYPE>PtrErr(func(v *bool) (bool, error) {
		if *v == false {
			return false, errors.New("false is invalid in this test")
		}
		return true, nil
	}, []*<TYPE>{&vt, &vf, &vf})
	if err == nil {
		t.Errorf("Remove<FTYPE>PtrErr failed.")
	}

	r, _ := Remove<FTYPE>PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("Remove<FTYPE>PtrErr failed.")
	}
}
`
}

// ReplaceActivityRemovePtrErr is template
func ReplaceActivityRemovePtrErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestRemoveIntPtrErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestRemoveIntPtrErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEven := func(num *string) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveStrPtrErr(partialIsEven, []*string{&v20, &v1, &v3, &v40})`
	s2 = `partialIsEven := func(num *string) (bool, error) {
		if *num == "0" {
			return false, errors.New("0 in invalid number for this test")
		}
		if *num == "20" || *num == "40" {
			return true, nil
		}
		return false, nil
	}
	NewList, _ = RemoveStrPtrErr(partialIsEven, []*string{&v20, &v1, &v3, &v40})`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEven := func(num *float32) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveFloat32PtrErr(partialIsEven, []*float32{&v20, &v1, &v3, &v40})`

	s2 = `partialIsEven := func(num *float32) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return int(*num)%10 == 0, nil
	}
	NewList, _ = RemoveFloat32PtrErr(partialIsEven, []*float32{&v20, &v1, &v3, &v40})`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEven := func(num *float64) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return *num%10 == 0, nil
	}
	NewList, _ = RemoveFloat64PtrErr(partialIsEven, []*float64{&v20, &v1, &v3, &v40})`

	s2 = `partialIsEven := func(num *float64) (bool, error) {
		if *num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return int(*num)%10 == 0, nil
	}
	NewList, _ = RemoveFloat64PtrErr(partialIsEven, []*float64{&v20, &v1, &v3, &v40})`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `_, err := RemoveStrPtrErr(func(v *string) (bool, error) {
		if *v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*string{&v20, &v1, &v3, &v40})`
	s2 = `_, err := RemoveStrPtrErr(func(v *string) (bool, error) {
		if *v == "20" {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []*string{&v20, &v1, &v3, &v40})`
	code = strings.Replace(code, s1, s2, -1)
	return code
}

//**********RemoveErr***********************

// RemoveErrTest is template
func RemoveErrTest() string {
	return `
func TestRemove<FTYPE>Err(t *testing.T) {
	// Test : even number in the list
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v20 <TYPE> = 20
	var v40 <TYPE> = 40

	expectedNewList := []<TYPE>{v1, v3}
	NewList, _ := Remove<FTYPE>Err(isEven<FTYPE>Err, []<TYPE>{v1, v2, v3, v4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("Remove<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []<TYPE>{v1, v3}
	partialIsEven := func(num <TYPE>) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = Remove<FTYPE>Err(partialIsEven, []<TYPE>{v20, v1, v3, v40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("Remove<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	r, _ := Remove<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Remove<FTYPE>Err failed.")
	}

	_, err := Remove<FTYPE>Err(func(v <TYPE>) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []<TYPE>{v20, v1, v3, v40})
	if err == nil {
		t.Errorf("Remove<FTYPE>Err failed.")
	}
}
`
}

// RemoveErrTestBool is template
func RemoveErrTestBool() string {
	return `
func TestRemove<FTYPE>Err(t *testing.T) {
	// Test : even number in the list
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedNewList := []<TYPE>{vt}
	NewList, _ := Remove<FTYPE>Err(func(v bool) (bool, error) { return v == false, nil }, []<TYPE>{vt, vf, vf})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("Remove<FTYPE>Err failed. Expected New list=%v, actual list=%v", expectedNewList[0], NewList[0])
	}

	_, err := Remove<FTYPE>Err(func(v bool) (bool, error) {
		if v == false {
			return false, errors.New("false is invalid in this test")
		}
		return true, nil
	}, []<TYPE>{vt, vf, vf})

	if err == nil {
		t.Errorf("Remove<FTYPE>Err failed.")
	}

	r, _ := Remove<FTYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Remove<FTYPE>Err failed.")
	}
}
`
}

// ReplaceActivityRemoveErr replaces ...
func ReplaceActivityRemoveErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestRemoveIntErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestRemoveIntErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEven := func(num string) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveStrErr(partialIsEven, []string{v20, v1, v3, v40})`
	s2 = `partialIsEven := func(num string) (bool, error) {
		if num == "0" {
			return false, errors.New("0 in invalid number for this test")
		}
		if num == "20" || num == "40" {
			return true, nil
		}
		return false, nil
	}
	NewList, _ = RemoveStrErr(partialIsEven, []string{v20, v1, v3, v40})`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEven := func(num float32) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveFloat32Err(partialIsEven, []float32{v20, v1, v3, v40})`

	s2 = `partialIsEven := func(num float32) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return int(num)%10 == 0, nil
	}
	NewList, _ = RemoveFloat32Err(partialIsEven, []float32{v20, v1, v3, v40})`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `partialIsEven := func(num float64) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return num%10 == 0, nil
	}
	NewList, _ = RemoveFloat64Err(partialIsEven, []float64{v20, v1, v3, v40})`

	s2 = `partialIsEven := func(num float64) (bool, error) {
		if num == 0 {
			return false, errors.New("0 in invalid number for this test")
		}
		return int(num)%10 == 0, nil
	}
	NewList, _ = RemoveFloat64Err(partialIsEven, []float64{v20, v1, v3, v40})`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `_, err := RemoveStrErr(func(v string) (bool, error) {
		if v == 20 {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []string{v20, v1, v3, v40})`
	s2 = `_, err := RemoveStrErr(func(v string) (bool, error) {
		if v == "20" {
			return false, errors.New("20 is invalid number for this test")
		}
		return true, nil
	}, []string{v20, v1, v3, v40})`
	code = strings.Replace(code, s1, s2, -1)
	return code
}
