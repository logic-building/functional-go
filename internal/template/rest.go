package template

// Rest is template to generate function(Rest) for user defined data type
func Rest() string {
	return `
func Rest<CONDITIONAL_TYPE>(l []<TYPE>) []<TYPE> {
	if l == nil {
		return []<TYPE>{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []<TYPE>{}
	}

	newList := make([]<TYPE>, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}
`
}

// RestPtr removes 1st item of the list and return new list having rest of the items
func RestPtr() string {
	return `

// Rest<CONDITIONAL_TYPE> removes 1st item of the list and return new list having rest of the items
func Rest<CONDITIONAL_TYPE>Ptr(l []*<TYPE>) []*<TYPE> {
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
