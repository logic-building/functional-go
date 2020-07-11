package basic

// PosWhtTest is template to generate itself for different combination of data type.
func PosWhtTest() string {
	return `
func TestPos<FTYPE>(t *testing.T) {
	r := Pos<FTYPE>Wht(1)
	if !r {
		t.Errorf("Pos<FTYPE>Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Pos<FTYPE>Wht(-1)
	if r {
		t.Errorf("Pos<FTYPE>Wht failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = 1
	rPtr := Pos<FTYPE>WhtPtr(&one)
	if !rPtr {
		t.Errorf("Pos<FTYPE>WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = Pos<FTYPE>WhtPtr(&zero)
	if rPtr {
		t.Errorf("Pos<FTYPE>WhtPtr failed. Expected=false, actual=true")
	}
}
`
}
