package basic

// ZeroPTest is template to generate itself for different combination of data type.
func ZeroPTest() string {
	return `
func TestZero<FTYPE>P(t *testing.T) {
	r := Zero<FTYPE>P(0)
	if !r {
		t.Errorf("Zero<FTYPE>P failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Zero<FTYPE>P(1)
	if r {
		t.Errorf("Zero<FTYPE>P failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = 1
	rPtr := Zero<FTYPE>PPtr(&zero)
	if !rPtr {
		t.Errorf("Zero<FTYPE>PPtr failed. Expected=true, actual=false")
	}

	rPtr = Zero<FTYPE>PPtr(&one)
	if rPtr {
		t.Errorf("Zero<FTYPE>PPtr failed. Expected=false, actual=true")
	}
}
`
}
