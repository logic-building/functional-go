package template

// Some is template to generate function(Some) for user defined data type
func Some() string {
	return `
func Some<CONDITIONAL_TYPE>(f func(<TYPE>) bool, list []<TYPE>) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
`
}

// SomePtr is template to generate function(Some) for user defined data type
func SomePtr() string {
	return `
func Some<CONDITIONAL_TYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
`
}
