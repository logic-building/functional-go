package basic

// ZeroTest is template to generate itself for different combination of data type.
func ZeroTest() string {
	return `
func TestZero<FTYPE>(t *testing.T) {
	r := Zero<FTYPE>(0)
	if !r {
		t.Errorf("Zero<FTYPE> failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Zero<FTYPE>(1)
	if r {
		t.Errorf("Zero<FTYPE> failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = 1
	rPtr := Zero<FTYPE>Ptr(&zero)
	if !rPtr {
		t.Errorf("Zero<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	rPtr = Zero<FTYPE>Ptr(&one)
	if rPtr {
		t.Errorf("Zero<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}
`
}
