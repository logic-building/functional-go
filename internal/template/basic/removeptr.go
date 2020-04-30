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

// RemovePtrErr removes the items from the given list based on supplied function and returns new list
func RemovePtrErr() string {
	return `
// Remove<FTYPE>PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *<TYPE> and return types(bool, error)
//	2. List of type: []*<TYPE>
//
// Returns:
//	New list and error: ([]*<TYPE>, error)
//	Empty list if both of arguments are nil or either one is nil.
func Remove<FTYPE>PtrErr(f func(*<TYPE>) (bool, error), list []*<TYPE>) ([]*<TYPE>, error) {
	if f == nil {
		return []*<TYPE>{}, nil
	}
	var newList []*<TYPE>
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}
`
}

// RemoveErr removes the items from the given list based on supplied function and returns new list
func RemoveErr() string {
	return `
// Remove<FTYPE>Err removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type <TYPE> and return types(bool, error)
//	2. List of type: []<TYPE>
//
// Returns:
//	New list and error: ([]<TYPE>, error)
//	Empty list if both of arguments are nil or either one is nil.
func Remove<FTYPE>Err(f func(<TYPE>) (bool, error), list []<TYPE>) ([]<TYPE>, error) {
	if f == nil {
		return []<TYPE>{}, nil
	}
	var newList []<TYPE>
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}
`
}
