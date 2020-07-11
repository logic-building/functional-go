package basic

// ZeroP is template
func ZeroP() string {
	return `
// Zero<FTYPE>P Returns true if num is zero, else false
func Zero<FTYPE>P(v <TYPE>) bool {
	if v == 0 {
		return true
	}
	return false
}

// Zero<FTYPE>PPtr Returns true if num is zero, else false
func Zero<FTYPE>PPtr(v *<TYPE>) bool {
	if *v == 0 {
		return true
	}
	return false
}
`
}
