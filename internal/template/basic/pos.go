package basic

// Pos is template
func Pos() string {
	return `
// Pos<FTYPE> Returns true if num is greater than zero, else false
func Pos<FTYPE>(v <TYPE>) bool {
	if v > 0 {
		return true
	}
	return false
}

// Pos<FTYPE>Ptr Returns true if num is greater than zero, else false
func Pos<FTYPE>Ptr(v *<TYPE>) bool {
	if *v > 0 {
		return true
	}
	return false
}
`
}
