package template

// TakeWhile is template to generate function(TakeWhile) for user defined data type
func TakeWhile() string {
	return `
func TakeWhile<CONDITIONAL_TYPE>(f func(<TYPE>) bool, list []<TYPE>) []<TYPE> {
	if f == nil {
		return []<TYPE>{}
	}
	var newList []<TYPE>
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}
`
}

// TakeWhilePtr is template to generate function(TakeWhile) for user defined data type
func TakeWhilePtr() string {
	return `
func TakeWhile<CONDITIONAL_TYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}
`
}
