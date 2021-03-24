package template

// Filter is template to generate function(Filter) for user defined data type
func Filter() string {
	return `
func Filter<CONDITIONAL_TYPE>(f func(<TYPE>) bool, list []<TYPE>) []<TYPE> {
	if f == nil {
		return []<TYPE>{}
	}
	var newList []<TYPE>
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
`
}

// FilterPtr is template to generate function(Filter) for user defined data type
func FilterPtr() string {
	return `
func Filter<CONDITIONAL_TYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
`
}
