package basic

// PosTest is template to generate itself for different combination of data type.
func NegTest() string {
	return `
func TestNeg<FTYPE>(t *testing.T) {
	r := Neg<FTYPE>(-1)
	if !r {
		t.Errorf("Neg<FTYPE> failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Neg<FTYPE>(1)
	if r {
		t.Errorf("Neg<FTYPE> failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = -1
	rPtr := Neg<FTYPE>Ptr(&one)
	if !rPtr {
		t.Errorf("Neg<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	rPtr = Neg<FTYPE>Ptr(&zero)
	if rPtr {
		t.Errorf("Neg<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}
`
}
