package basic

// Zero is template
func Zero() string {
	return `
// Zero<FTYPE> Returns true if num is zero, else false
func Zero<FTYPE>(v <TYPE>) bool {
	if v == 0 {
		return true
	}
	return false
}

// Zero<FTYPE>Ptr Returns true if num is zero, else false
func Zero<FTYPE>Ptr(v *<TYPE>) bool {
	if *v == 0 {
		return true
	}
	return false
}
`
}
