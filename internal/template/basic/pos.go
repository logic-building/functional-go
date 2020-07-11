package basic

// PosWht is template
func PosWht() string {
	return `
// PosWht<FTYPE> Returns true if num is greater than zero, else false
func PosWht<FTYPE>(v <TYPE>) bool {
	if v > 0 {
		return true
	}
	return false
}

// PosWht<FTYPE>Ptr Returns true if num is greater than zero, else false
func PosWht<FTYPE>Ptr(v *<TYPE>) bool {
	if *v > 0 {
		return true
	}
	return false
}
`
}
