package basic

// OddP is template
func OddP() string {
	return `
// Odd<FTYPE>P Returns true if n is odd
func Odd<FTYPE>P(v <TYPE>) bool {
	return v%2 != 0
}

// Odd<FTYPE>PPtr Returns true if n is odd
func Odd<FTYPE>PPtr(v *<TYPE>) bool {
	return *v%2 != 0
}
`
}
