package basic

// ExistsPtr is template
func ExistsPtr() string {
	return `
// Exists<FTYPE>Ptr checks if given item exists in the list
func Exists<FTYPE>Ptr(num *<TYPE>, list []*<TYPE>) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}
`
}
