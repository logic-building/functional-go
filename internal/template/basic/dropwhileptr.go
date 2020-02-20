package basic

// Drop returns a new list after dropping the given item
func DropWhilePtr() string {
	return `
/// DropWhile<FTYPE>Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil

func DropWhile<FTYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*<TYPE>, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}
`
}
