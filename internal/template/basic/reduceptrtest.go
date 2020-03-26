package basic

// Reduce<FTYPE>Ptr reduces a list to a single value by combining elements via a supplied function
func ReducePtrTest() string {
	return `
func TestReduce<FTYPE>Ptr(t *testing.T) {
	var v1 <TYPE> = 1
	var v2 <TYPE> = 2
	var v3 <TYPE> = 3
	var v4 <TYPE> = 4
	var v5 <TYPE> = 5

	list := []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	var expected <TYPE> = 15
	actual := Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list, 3)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{&v1, &v2}
	expected = 3
	actual = Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{&v1}
	expected = 1
	actual = Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>Ptr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*<TYPE>{}
	expected = 0
	actual = Reduce<FTYPE>Ptr(plus<FTYPE>Ptr, list)
	if *actual != expected {
		t.Errorf("Reduce<FTYPE>Ptr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}

func plus<FTYPE>Ptr(num1, num2 *<TYPE>) *<TYPE> {
	c := *num1 + *num2
	return &c
}

`
}
