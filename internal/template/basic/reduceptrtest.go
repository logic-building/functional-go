package basic

import "strings"

// ReducePtrTest reduces a list to a single value by combining elements via a supplied function
func ReducePtrTest() string {
	return `
func TestReduce<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	list := []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	var expected <TYPE> = 15
	actual := Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list, 3)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{&v1, &v2}
	expected = 3
	actual = Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{&v1}
	expected = 1
	actual = Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{}
	expected = 0
	actual = Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plus<FTYPE>Ptr(num1, num2 *<TYPE>) *<TYPE> {
	c := *num1 + *num2
	return &c
}
`
}

// **********Reduce<FTYPE>PtrErr*************

// ReducePtrErrTest reduces a list to a single value by combining elements via a supplied function
func ReducePtrErrTest() string {
	return `
func TestReduce<FTYPE>PtrErr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v0 <TYPE>

	list := []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	var expected <TYPE> = 15
	actual, _ := Reduce<FTYPE>PtrErr(plus<FTYPE>PtrErr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>PtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list2 := []*<TYPE>{&v1, &v0, &v3, &v4, &v5}
	_, err := Reduce<FTYPE>PtrErr(plus<FTYPE>PtrErr, list2)
	if err == nil {
		t.Errorf("Reduce<FTYPE>PtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual, _ = Reduce<FTYPE>PtrErr(plus<FTYPE>PtrErr, list, 3)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>PtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{&v1, &v2}
	expected = 3
	actual, _ = Reduce<FTYPE>PtrErr(plus<FTYPE>PtrErr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>PtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{&v1}
	expected = 1
	actual, _ = Reduce<FTYPE>PtrErr(plus<FTYPE>PtrErr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>PtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{}
	expected = 0
	actual, _ = Reduce<FTYPE>PtrErr(plus<FTYPE>PtrErr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>PtrErr failed. actual=%v, expected=%v", *actual, expected)
	}
}

func plus<FTYPE>PtrErr(num1, num2 *<TYPE>) (*<TYPE>, error) {
	if *num1 == 0 || *num2 == 0 {
		return nil, errors.New("0 in not valid number for this test")
	}
	c := *num1 + *num2
	return &c, nil
}
`
}

// ReplaceActivityReducePtrErr replaces ...
func ReplaceActivityReducePtrErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestReduceIntPtrErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestReduceIntPtrErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestReduceStrPtrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	var v0 string

	list := []*string{&v1, &v2, &v3, &v4, &v5}
	var expected string = 15
	actual, _ := ReduceStrPtrErr(plusStrPtrErr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list2 := []*string{&v1, &v0, &v3, &v4, &v5}
	_, err := ReduceStrPtrErr(plusStrPtrErr, list2)
	if err == nil {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual, _ = ReduceStrPtrErr(plusStrPtrErr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2}
	expected = 3
	actual, _ = ReduceStrPtrErr(plusStrPtrErr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1}
	expected = 1
	actual, _ = ReduceStrPtrErr(plusStrPtrErr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{}
	expected = 0
	actual, _ = ReduceStrPtrErr(plusStrPtrErr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}
}

func plusStrPtrErr(num1, num2 *string) (*string, error) {
	if *num1 == 0 || *num2 == 0 {
		return nil, errors.New("0 in not valid number for this test")
	}
	c := *num1 + *num2
	return &c, nil
}`
	s2 = `func TestReduceStrPtrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	var v0 string = "0"

	list := []*string{&v1, &v2, &v3, &v4, &v5}
	var expected string = "12345"
	actual, _ := ReduceStrPtrErr(plusStrPtrErr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list2 := []*string{&v1, &v0, &v3, &v4, &v5}
	_, err := ReduceStrPtrErr(plusStrPtrErr, list2)
	if err == nil {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2, &v3, &v4, &v5}
	expected = "312345"
	actual, _ = ReduceStrPtrErr(plusStrPtrErr, list, "3")
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2}
	expected = "12"
	actual, _ = ReduceStrPtrErr(plusStrPtrErr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1}
	expected = "1"
	actual, _ = ReduceStrPtrErr(plusStrPtrErr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{}
	expected = ""
	actual, _ = ReduceStrPtrErr(plusStrPtrErr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtrErr failed. actual=%v, expected=%v", *actual, expected)
	}
}

func plusStrPtrErr(num1, num2 *string) (*string, error) {
	if *num1 == "0" || *num2 == "0" {
		return nil, errors.New("0 in not valid number for this test")
	}
	c := *num1 + *num2
	return &c, nil
}`

	code = strings.Replace(code, s1, s2, -1)

	return code
}

// **********Reduce<FTYPE>Err*************

// ReduceErrTest reduces a list to a single value by combining elements via a supplied function
func ReduceErrTest() string {
	return `
func TestReduce<FTYPE>Err(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5
	var v0 <TYPE> = 0

	list := []<TYPE>{v1, v2, v3, v4, v5}
	var expected <TYPE> = 15
	actual, _ := Reduce<FTYPE>Err(plus<FTYPE>Err, list)
	if actual != expected {
		t.Errorf("Reduce<FTYPE>Err failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []<TYPE>{v1, v0, v3, v4, v5}
	_, err := Reduce<FTYPE>Err(plus<FTYPE>Err, list2)
	if err == nil {
		t.Errorf("Reduce<FTYPE>Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []<TYPE>{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = Reduce<FTYPE>Err(plus<FTYPE>Err, list, 3)
	if actual != expected {
		t.Errorf("Reduce<FTYPE>Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []<TYPE>{v1, v2}
	expected = 3
	actual, _ = Reduce<FTYPE>Err(plus<FTYPE>Err, list)
	if actual != expected {
		t.Errorf("Reduce<FTYPE>Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []<TYPE>{v1}
	expected = 1
	actual, _ = Reduce<FTYPE>Err(plus<FTYPE>Err, list)
	if actual != expected {
		t.Errorf("Reduce<FTYPE>Err failed. actual=%v, expected=%v", actual, expected)
	}

	list = []<TYPE>{}
	expected = 0
	actual, _ = Reduce<FTYPE>Err(plus<FTYPE>Err, list)
	if actual != expected {
		t.Errorf("Reduce<FTYPE>Err failed. actual=%v, expected=%v", actual, expected)
	}
}

func plus<FTYPE>Err(num1, num2 <TYPE>) (<TYPE>, error) {
	if num1 == 0 || num2 == 0 {
		return <TYPE>(0), errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}
`
}

// ReplaceActivityReduceErr replaces ...
func ReplaceActivityReduceErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestReduceIntErr(t *testing.T) {`
	s2 := `import (
	"errors"
	"testing"
)

func TestReduceIntErr(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestReduceStrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	var v0 string = "0"

	list := []string{v1, v2, v3, v4, v5}
	var expected string = 15
	actual, _ := ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []string{v1, v0, v3, v4, v5}
	_, err := ReduceStrErr(plusStrErr, list2)
	if err == nil {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{v1, v2, v3, v4, v5}
	expected = 18
	actual, _ = ReduceStrErr(plusStrErr, list, 3)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{v1, v2}
	expected = 3
	actual, _ = ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{v1}
	expected = 1
	actual, _ = ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{}
	expected = 0
	actual, _ = ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusStrErr(num1, num2 string) (string, error) {
	if num1 == 0 || num2 == 0 {
		return 0, errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}`
	s2 = `func TestReduceStrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	var v0 string = "0"

	list := []string{v1, v2, v3, v4, v5}
	var expected string = "12345"
	actual, _ := ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list2 := []string{v1, v0, v3, v4, v5}
	_, err := ReduceStrErr(plusStrErr, list2)
	if err == nil {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{v1, v2, v3, v4, v5}
	expected = "312345"
	actual, _ = ReduceStrErr(plusStrErr, list, "3")
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{v1, v2}
	expected = "12"
	actual, _ = ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{v1}
	expected = "1"
	actual, _ = ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{}
	expected = ""
	actual, _ = ReduceStrErr(plusStrErr, list)
	if actual != expected {
		t.Errorf("ReduceStrErr failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusStrErr(num1, num2 string) (string, error) {
	if num1 == "0" || num2 == "0" {
		return "", errors.New("0 in not valid number for this test")
	}
	c := num1 + num2
	return c, nil
}`

	code = strings.Replace(code, s1, s2, -1)

	return code
}
