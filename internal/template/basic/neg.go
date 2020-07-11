package basic

// NegWht is template
func NegWht() string {
	return `
// Neg<FTYPE>Wht Returns true if num is less than zero, else false
func Neg<FTYPE>Wht(v <TYPE>) bool {
	if v < 0 {
		return true
	}
	return false
}

// Neg<FTYPE>WhtPtr Returns true if num is less than zero, else false
func Neg<FTYPE>WhtPtr(v *<TYPE>) bool {
	if *v < 0 {
		return true
	}
	return false
}
`
}
