package basic

// RestPtr removes 1st item of the list and return new list having rest of the items
func RestPtr() string {
	return `
// Rest<FTYPE>Ptr removes 1st item of the list and return new list having rest of the items
func Rest<FTYPE>Ptr(l []*<TYPE>) []*<TYPE> {
	if l == nil {
		return []*<TYPE>{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*<TYPE>{}
	}

	newList := make([]*<TYPE>, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}
`
}
