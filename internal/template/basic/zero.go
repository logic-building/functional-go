package basic

// ZeroWht is template
func ZeroWht() string {
	return `
// Zero<FTYPE>Wht Returns true if num is zero, else false
func Zero<FTYPE>Wht(v <TYPE>) bool {
	if v == 0 {
		return true
	}
	return false
}

// Zero<FTYPE>WhtPtr Returns true if num is zero, else false
func Zero<FTYPE>WhtPtr(v *<TYPE>) bool {
	if *v == 0 {
		return true
	}
	return false
}
`
}
