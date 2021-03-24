package basic

// DropWhilePtr returns a new list after dropping the given item
func DropWhilePtr() string {
	return `
// DropWhile<FTYPE>Ptr drops the items from the list as long as condition satisfies.
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

// DropWhilePtrErr returns a new list after dropping the given item
func DropWhilePtrErr() string {
	return `
// DropWhile<FTYPE>PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhile<FTYPE>PtrErr(f func(*<TYPE>) (bool, error), list []*<TYPE>) ([]*<TYPE>, error) {
	if f == nil {
		return []*<TYPE>{}, nil
	}
	var newList []*<TYPE>
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*<TYPE>, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}
`
}

// DropWhileErr returns a new list after dropping the given item
func DropWhileErr() string {
	return `
// DropWhile<FTYPE>Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhile<FTYPE>Err(f func(<TYPE>) (bool, error), list []<TYPE>) ([]<TYPE>, error) {
	if f == nil {
		return []<TYPE>{}, nil
	}
	var newList []<TYPE>
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]<TYPE>, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}
`
}
