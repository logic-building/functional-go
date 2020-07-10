package basic

// Odd is template
func Odd() string {
	return `
// Odd<FTYPE> Returns true if n is odd
func Odd<FTYPE>(v <TYPE>) bool {
	return v%2 != 0
}

// Odd<FTYPE>Ptr Returns true if n is odd
func Odd<FTYPE>Ptr(v *<TYPE>) bool {
	return *v%2 != 0
}
`
}
