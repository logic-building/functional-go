package basic

// PosWht is template
func PosWht() string {
	return `
// Pos<FTYPE>Wht Returns true if num is greater than zero, else false
func Pos<FTYPE>Wht(v <TYPE>) bool {
	if v > 0 {
		return true
	}
	return false
}

// Pos<FTYPE>WhtPtr Returns true if num is greater than zero, else false
func Pos<FTYPE>WhtPtr(v *<TYPE>) bool {
	if *v > 0 {
		return true
	}
	return false
}
`
}
