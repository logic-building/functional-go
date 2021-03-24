package basic

// EvenPTest is template to generate itself for different combination of data type.
func EvenPTest() string {
	return `
func TestEven<FTYPE>(t *testing.T) {
	r := Even<FTYPE>P(10)
	if !r {
		t.Errorf("Even<FTYPE>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Even<FTYPE>P(1)
	if r {
		t.Errorf("Even<FTYPE>P failed. Expected=false, actual=true")
	}

	var two <TYPE> = 2
	rPtr := Even<FTYPE>PPtr(&two)
	if !rPtr {
		t.Errorf("Even<FTYPE>PPtr failed. Expected=true, actual=false")
	}
}
`
}
