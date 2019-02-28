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
