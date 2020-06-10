package basic

// MaxPtrTest is template
func MaxPtrTest() string {
	return `
func TestMax<FTYPE>Ptr(t *testing.T) {
	// Test : get max number in the list
	var v2 <TYPE> = 2
	var v4 <TYPE> = 4
	var v8 <TYPE> = 8
	var v10 <TYPE> = 10

	list := []*<TYPE>{&v8, &v2, &v10, &v4}
	max := Max<FTYPE>Ptr(list)
	if *max != 10 {
		t.Errorf("Max<FTYPE>Ptr failed. Expected=10, actual=%v", max)
		t.Errorf(reflect.String.String())
	}
}
`
}
