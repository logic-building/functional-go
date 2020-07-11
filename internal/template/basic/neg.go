package basic

// NegP is template
func NegP() string {
	return `
// Neg<FTYPE>P Returns true if num is less than zero, else false
func Neg<FTYPE>P(v <TYPE>) bool {
	if v < 0 {
		return true
	}
	return false
}

// Neg<FTYPE>PPtr Returns true if num is less than zero, else false
func Neg<FTYPE>PPtr(v *<TYPE>) bool {
	if *v < 0 {
		return true
	}
	return false
}
`
}
