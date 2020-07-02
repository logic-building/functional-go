package basic

// PosTest is template to generate itself for different combination of data type.
func PosTest() string {
	return `
func TestPos<FTYPE>(t *testing.T) {
	r := Pos<FTYPE>(1)
	if !r {
		t.Errorf("Pos<FTYPE> failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Pos<FTYPE>(-1)
	if r {
		t.Errorf("Pos<FTYPE> failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = 1
	rPtr := Pos<FTYPE>Ptr(&one)
	if !rPtr {
		t.Errorf("Pos<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	rPtr = Pos<FTYPE>Ptr(&zero)
	if rPtr {
		t.Errorf("Pos<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}
`
}
