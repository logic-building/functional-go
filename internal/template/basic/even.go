package basic

// Even is template
func Even() string {
	return `
// Even<FTYPE> Returns true if n is even
func Even<FTYPE>(v <TYPE>) bool {
	return v%2 == 0
}

// Even<FTYPE>Ptr Returns true if n is even
func Even<FTYPE>Ptr(v *<TYPE>) bool {
	return *v%2 == 0
}
`
}
