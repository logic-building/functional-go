package basic

// NegPTest is template to generate itself for different combination of data type.
func NegPTest() string {
	return `
func TestNeg<FTYPE>P(t *testing.T) {
	r := Neg<FTYPE>P(-1)
	if !r {
		t.Errorf("Neg<FTYPE>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Neg<FTYPE>P(1)
	if r {
		t.Errorf("Neg<FTYPE>P failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = -1
	rPtr := Neg<FTYPE>PPtr(&one)
	if !rPtr {
		t.Errorf("Neg<FTYPE>PPtr failed. Expected=true, actual=false")
	}

	rPtr = Neg<FTYPE>PPtr(&zero)
	if rPtr {
		t.Errorf("Neg<FTYPE>PPtr failed. Expected=false, actual=true")
	}
}
`
}
