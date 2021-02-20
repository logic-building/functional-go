package template

// Every is template to generate function(Every) for user defined data type
func Every() string {
	return `
func Every<CONDITIONAL_TYPE>(f func(<TYPE>) bool, list []<TYPE>) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
`
}

// EveryPtr is template to generate function(Every) for user defined data type
func EveryPtr() string {
	return `
func Every<CONDITIONAL_TYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
`
}
