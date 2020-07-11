package basic

// NegWhtTest is template to generate itself for different combination of data type.
func NegWhtTest() string {
	return `
func TestNeg<FTYPE>Wht(t *testing.T) {
	r := Neg<FTYPE>Wht(-1)
	if !r {
		t.Errorf("Neg<FTYPE>Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Neg<FTYPE>Wht(1)
	if r {
		t.Errorf("Neg<FTYPE>Wht failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = -1
	rPtr := Neg<FTYPE>WhtPtr(&one)
	if !rPtr {
		t.Errorf("Neg<FTYPE>WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = Neg<FTYPE>WhtPtr(&zero)
	if rPtr {
		t.Errorf("Neg<FTYPE>WhtPtr failed. Expected=false, actual=true")
	}
}
`
}
