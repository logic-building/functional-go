package basic

import "strings"

// MapIONumber is template to generate itself for different combination of data type.
func MapIONumber() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []<OUTPUT_TYPE>{2, 3, 4}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrNumber is template to generate itself for different combination of data type.
func MapIOStrNumber() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIONumberStr is template to generate itself for different combination of data type.
func MapIONumberStr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10"}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{10})

	if newList[0] != expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIONumberBool is template to generate itself for different combination of data type.
func MapIONumberBool() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrBool is template to generate itself for different combination of data type.
func MapIOStrBool() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{"10", "0"})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolNumber is template to generate itself for different combination of data type.
func MapIOBoolNumber() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10, 0}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolStr is template to generate itself for different combination of data type.
func MapIOBoolStr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10", "0"}
	newList := Map<FINPUT_TYPE><FOUTPUT_TYPE>(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>, []<INPUT_TYPE>{true, false})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, nil)) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	if len(Map<FINPUT_TYPE><FOUTPUT_TYPE>(nil, []<INPUT_TYPE>{})) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIONumberErr is template to generate itself for different combination of data type.
func MapIONumberErr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []<OUTPUT_TYPE>{2, 3, 6}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{1, 2, 5})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(plusOne<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{1, 2, 3})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrNumberErr is template to generate itself for different combination of data type.
func MapIOStrNumberErr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"0"})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (<OUTPUT_TYPE>, error) {
	if num == "0" {
		return 0, errors.New("0 is invalid number for this test")
	}
	if num == "ten" {
		return <OUTPUT_TYPE>(10), nil
	}
	return 0, nil
}
`
}

// MapIONumberStrErr is template to generate itself for different combination of data type.
func MapIONumberStrErr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10"}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10})

	if newList[0] != expectedList[0] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{0})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}

func someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err(num <INPUT_TYPE>) (<OUTPUT_TYPE>, error) {
	if num == 0 {
		return "", errors.New("0 is not valid number for this test")
	}
	if num == 10 {
		return string("10"), nil
	}
	return "0", nil
}
`
}

// MapIONumberBoolErr is template to generate itself for different combination of data type.
func MapIONumberBoolErr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 0})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{10, 3})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOStrBoolErr is template to generate itself for different combination of data type.
func MapIOStrBoolErr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{true, false}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "0"})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{"10", "3"})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolNumberErr is template to generate itself for different combination of data type.
func MapIOBoolNumberErr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{10, 0}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE> failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, false})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// MapIOBoolStrErr is template to generate itself for different combination of data type.
func MapIOBoolStrErr() string {
	return `
func TestMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(t *testing.T) {
	// Test : someLogic
	expectedList := []<OUTPUT_TYPE>{"10", "0"}
	newList, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, true})

	if newList[0] != expectedList[0] && newList[1] != expectedList[1] {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}

	r, _ = Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(nil, []<INPUT_TYPE>{})
	if len(r) > 0 {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	_, err := Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(someLogic<FINPUT_TYPE><FOUTPUT_TYPE>Err, []<INPUT_TYPE>{true, false})
	if err == nil {
		t.Errorf("Map<FINPUT_TYPE><FOUTPUT_TYPE>Err failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
}
`
}

// ReplaceActivityMapIOErr replaces ...
func ReplaceActivityMapIOErr(code string) string {
	s1 := `import (
	_ "errors"
	"reflect"
	"testing"
)

func TestMapIntInt64Err(t *testing.T) {`
	s2 := `import (
	"errors"
	"reflect"
	"testing"
)

func TestMapIntInt64Err(t *testing.T) {`

	code = strings.Replace(code, s1, s2, -1)
	return code
}
