package basic

// PosP is template
func PosP() string {
	return `
// Pos<FTYPE>P Returns true if num is greater than zero, else false
func Pos<FTYPE>P(v <TYPE>) bool {
	if v > 0 {
		return true
	}
	return false
}

// Pos<FTYPE>PPtr Returns true if num is greater than zero, else false
func Pos<FTYPE>PPtr(v *<TYPE>) bool {
	if *v > 0 {
		return true
	}
	return false
}
`
}
