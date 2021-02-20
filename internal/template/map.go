package template

// Map is template to generate function(Map) for user defined data type
func Map() string {
	return `
func Map<CONDITIONAL_TYPE>(f func(<TYPE>) <TYPE>, list []<TYPE>) []<TYPE> {
	if f == nil {
		return []<TYPE>{}
	}
	newList := make([]<TYPE>, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
`
}

// MapPtr is template
func MapPtr() string {
	return `
func Map<CONDITIONAL_TYPE>Ptr(f func(*<TYPE>) *<TYPE>, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	newList := make([]*<TYPE>, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
`
}
