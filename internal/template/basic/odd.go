package basic

// OddWht is template
func OddWht() string {
	return `
// Odd<FTYPE>Wht Returns true if n is odd
func Odd<FTYPE>Wht(v <TYPE>) bool {
	return v%2 != 0
}

// Odd<FTYPE>WhtPtr Returns true if n is odd
func Odd<FTYPE>WhtPtr(v *<TYPE>) bool {
	return *v%2 != 0
}
`
}
