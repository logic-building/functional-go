package basic

// OddPTest is template to generate itself for different combination of data type.
func OddPTest() string {
	return `
func TestOdd<FTYPE>(t *testing.T) {
	r := Odd<FTYPE>P(11)
	if !r {
		t.Errorf("Odd<FTYPE>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Odd<FTYPE>P(2)
	if r {
		t.Errorf("Odd<FTYPE>P failed. Expected=false, actual=true")
	}

	var three <TYPE> = 3
	rPtr := Odd<FTYPE>PPtr(&three)
	if !rPtr {
		t.Errorf("Odd<FTYPE>PPtr failed. Expected=true, actual=false")
	}
}
`
}
