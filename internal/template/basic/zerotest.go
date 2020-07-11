package basic

// ZeroWhtTest is template to generate itself for different combination of data type.
func ZeroWhtTest() string {
	return `
func TestZero<FTYPE>Wht(t *testing.T) {
	r := Zero<FTYPE>Wht(0)
	if !r {
		t.Errorf("Zero<FTYPE>Wht failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Zero<FTYPE>Wht(1)
	if r {
		t.Errorf("Zero<FTYPE>Wht failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = 1
	rPtr := Zero<FTYPE>WhtPtr(&zero)
	if !rPtr {
		t.Errorf("Zero<FTYPE>WhtPtr failed. Expected=true, actual=false")
	}

	rPtr = Zero<FTYPE>WhtPtr(&one)
	if rPtr {
		t.Errorf("Zero<FTYPE>WhtPtr failed. Expected=false, actual=true")
	}
}
`
}
