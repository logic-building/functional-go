package basic

// EvenTest is template to generate itself for different combination of data type.
func EvenTest() string {
	return `
func TestEven<FTYPE>(t *testing.T) {
	r := Even<FTYPE>(10)
	if !r {
		t.Errorf("Even<FTYPE> failed. Expected=true, actual=false")
		t.Errorf(reflect.String.String())
	}

	r = Even<FTYPE>(1)
	if r {
		t.Errorf("Even<FTYPE> failed. Expected=false, actual=true")
	}

	var two <TYPE> = 2
	rPtr := Even<FTYPE>Ptr(&two)
	if !rPtr {
		t.Errorf("Even<FTYPE>Ptr failed. Expected=true, actual=false")
	}
}
`
}
