package basic

// DedupeTest is template to generate itself for different combination of data type.
func DedupeTest() string {
	return `
func TestDedupe<FTYPE>(t *testing.T) {
	var v0 <TYPE>
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4

	expectedList := []<TYPE>{v0, v1, v2, v3, v4, v3}
	givenList := []<TYPE>{v0, v0, v1, v2, v2, v3, v3, v3, v4, v3, v3}
	r := Dedupe<FTYPE>(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupe<FTYPE> failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*<TYPE>{&v0, &v1, &v2, &v3, &v4, &v3}
	givenListPtr := []*<TYPE>{&v0, &v0, &v1, &v2, &v2, &v3, &v3, &v3, &v4, &v3, &v3}
	rPtr := Dedupe<FTYPE>Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupe<FTYPE> failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupe<FTYPE>Ptr failed.")
	}
}
`
}

// DedupeBoolTest is template to generate itself for different combination of data type.
func DedupeBoolTest() string {
	return `
func TestDedupe<FTYPE>(t *testing.T) {
	var vt <TYPE> = true
	var vf <TYPE> = false

	expectedList := []<TYPE>{vt, vf, vt}
	givenList := []<TYPE>{vt, vt, vf, vf, vf, vt, vt, vt, vt, vt, vt}
	r := Dedupe<FTYPE>(givenList)
	if !reflect.DeepEqual(r, expectedList) {
		t.Errorf("TestDedupe<FTYPE> failed. acutal_list=%v, expected_list=%v", r, expectedList)
	}

	expectedListPtr := []*<TYPE>{&vt, &vf, &vt, &vf}
	givenListPtr := []*<TYPE>{&vt, &vt, &vf, &vf, &vf, &vt, &vf, &vf, &vf, &vf, &vf}
	rPtr := Dedupe<FTYPE>Ptr(givenListPtr)
	if !reflect.DeepEqual(rPtr, expectedListPtr) {
		t.Errorf("TestDedupe<FTYPE>Ptr failed. acutal_list=%v, expected_listPtr=%v", rPtr, expectedListPtr)
	}

	if *expectedListPtr[0] != *rPtr[0] {
		t.Errorf("TestDedupe<FTYPE>Ptr failed.")
	}
}
`
}
