package basic

// Neg is template
func Neg() string {
	return `
// Neg<FTYPE> Returns true if num is less than zero, else false
func Neg<FTYPE>(v <TYPE>) bool {
	if v < 0 {
		return true
	}
	return false
}

// Neg<FTYPE>Ptr Returns true if num is less than zero, else false
func Neg<FTYPE>Ptr(v *<TYPE>) bool {
	if *v < 0 {
		return true
	}
	return false
}
`
}
