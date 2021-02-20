package methodchain

import "strings"

// MethodChainMapTest is template to generate itself for different combination of data type.
func MethodChainMapTest() string {
	return `
func TestMap<FTYPE>MethodChain(t *testing.T) {
	expectedSquareList := []<TYPE>{1, 4, 9}
	squareList := Make<FTYPE>Slice([]<TYPE>{1, 2, 3}...).Map(square<FTYPE>)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("Map<FTYPE>MthodChain failed")
	}

	if len(Make<FTYPE>Slice().Map(square<FTYPE>)) > 0 {
		t.Errorf("Map<FTYPE> failed.")
		t.Errorf(reflect.String.String())
	}
}
`
}

// MethodChainMapBoolTest is template to generate itself for different combination of data type.
func MethodChainMapBoolTest() string {
	return `
func TestMap<FTYPE>MethodChain(t *testing.T) {
	expectedSquareList := []<TYPE>{false, true, false}
	squareList := Make<FTYPE>Slice([]<TYPE>{true, false, true}...).Map(inverseBool)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("Map<FTYPE>MthodChain failed")
	}

	if len(Make<FTYPE>Slice().Map(inverseBool)) > 0 {
		t.Errorf("Map<FTYPE> failed.")
		t.Errorf(reflect.String.String())
	}
}

func inverseBool(v bool) bool {
	if v == true {
		return false
	} 
	return true
}
`
}

// ReplaceActivityMethodChainMap - Replace activity for string type
func ReplaceActivityMethodChainMap(code string) string {
	t1 := `func TestMapStrMethodChain(t *testing.T) {
	expectedSquareList := []string{1, 4, 9}
	squareList := MakeStrSlice([]string{1, 2, 3}...).Map(squareStr)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapStrMthodChain failed")
	}

	if len(MakeStrSlice().Map(squareStr)) > 0 {
		t.Errorf("MapStr failed.")
		t.Errorf(reflect.String.String())
	}
}`
	t2 := `func TestMapStrMethodChain(t *testing.T) {
	expectedSquareList := []string{"11", "22", "33"}
	squareList := MakeStrSlice([]string{"1", "2", "3"}...).Map(squareStr)

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapStrMthodChain failed")
	}

	if len(MakeStrSlice().Map(squareStr)) > 0 {
		t.Errorf("MapStr failed.")
		t.Errorf(reflect.String.String())
	}
}

func squareStr(s string) string {
	return s+s
}`
	code = strings.ReplaceAll(code, t1, t2)
	return code
}
