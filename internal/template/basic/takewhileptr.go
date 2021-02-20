package basic

// TakeWhilePtr is template to generate function(TakeWhile) for user defined data type
func TakeWhilePtr() string {
	return `
// TakeWhile<FTYPE>Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhile<FTYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}
`
}

// TakeWhilePtrErr is template to generate function(TakeWhile) for user defined data type
func TakeWhilePtrErr() string {
	return `
// TakeWhile<FTYPE>PtrErr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type *<TYPE> and returns ([]*<TYPE>, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhile<FTYPE>PtrErr(f func(*<TYPE>) (bool, error), list []*<TYPE>) ([]*<TYPE>, error) {
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
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}
`
}

// TakeWhileErr is template to generate function(TakeWhile) for user defined data type
func TakeWhileErr() string {
	return `
// TakeWhile<FTYPE>Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type <TYPE> and returns ([]<TYPE>, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhile<FTYPE>Err(f func(<TYPE>) (bool, error), list []<TYPE>) ([]<TYPE>, error) {
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
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}
`
}
