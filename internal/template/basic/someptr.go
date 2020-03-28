package basic

// Some is template to generate function(Some) for user defined data type
func SomePtr() string {
	return `
// Some<FTYPE>Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func Some<FTYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
`
}
