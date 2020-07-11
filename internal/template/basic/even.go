package basic

// EvenP is template
func EvenP() string {
	return `
// Even<FTYPE>P Returns true if n is even
func Even<FTYPE>P(v <TYPE>) bool {
	return v%2 == 0
}

// Even<FTYPE>PPtr Returns true if n is even
func Even<FTYPE>PPtr(v *<TYPE>) bool {
	return *v%2 == 0
}
`
}
