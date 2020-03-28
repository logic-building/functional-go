package basic

// RemovePtr removes the items from the given list based on supplied function and returns new list
func RemovePtr() string {
	return `
// Remove<FTYPE>Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func Remove<FTYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
`
}
