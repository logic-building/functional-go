package basic

// PosPTest is template to generate itself for different combination of data type.
func PosPTest() string {
	return `
func TestPos<FTYPE>(t *testing.T) {
	r := Pos<FTYPE>P(1)
	if !r {
		t.Errorf("Pos<FTYPE>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Pos<FTYPE>P(-1)
	if r {
		t.Errorf("Pos<FTYPE>P failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = 1
	rPtr := Pos<FTYPE>PPtr(&one)
	if !rPtr {
		t.Errorf("Pos<FTYPE>PPtr failed. Expected=true, actual=false")
	}

	rPtr = Pos<FTYPE>PPtr(&zero)
	if rPtr {
		t.Errorf("Pos<FTYPE>PPtr failed. Expected=false, actual=true")
	}
}
`
}
