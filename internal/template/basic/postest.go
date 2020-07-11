package basic

// PosWhtTest is template to generate itself for different combination of data type.
func PosWhtTest() string {
	return `
func TestPos<FTYPE>(t *testing.T) {
	r := PosWht<FTYPE>(1)
	if !r {
		t.Errorf("PosWht<FTYPE> failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = PosWht<FTYPE>(-1)
	if r {
		t.Errorf("PosWht<FTYPE> failed. Expected=false, actual=true")
	}

	var zero <TYPE>
	var one <TYPE> = 1
	rPtr := PosWht<FTYPE>Ptr(&one)
	if !rPtr {
		t.Errorf("PosWht<FTYPE>Ptr failed. Expected=true, actual=false")
	}

	rPtr = PosWht<FTYPE>Ptr(&zero)
	if rPtr {
		t.Errorf("PosWht<FTYPE>Ptr failed. Expected=false, actual=true")
	}
}
`
}
