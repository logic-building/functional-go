package basic

// OddTest is template to generate itself for different combination of data type.
func OddTest() string {
	return `
func TestOdd<FTYPE>(t *testing.T) {
	r := Odd<FTYPE>Wht(11)
	if !r {
		t.Errorf("Odd<FTYPE>Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Odd<FTYPE>Wht(2)
	if r {
		t.Errorf("Odd<FTYPE>Wht failed. Expected=false, actual=true")
	}

	var three <TYPE> = 3
	rPtr := Odd<FTYPE>WhtPtr(&three)
	if !rPtr {
		t.Errorf("Odd<FTYPE>WhtPtr failed. Expected=true, actual=false")
	}
}
`
}
