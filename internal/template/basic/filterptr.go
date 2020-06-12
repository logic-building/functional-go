package basic

// FilterPtr filters list based on function passed as 1st argument
func FilterPtr() string {
	return `
// Filter<FTYPE>Ptr takes two arguments
//  1. Function: takes 1 argument of type <TYPE> and returns bool
//  2. slice of type []*<TYPE>
//
// Returns:
//  new filtered list
func Filter<FTYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
`
}

// FilterPtrErr filters list based on function passed as 1st argument
func FilterPtrErr() string {
	return `
// Filter<FTYPE>PtrErr takes two arguments
//  1. Function: takes 1 argument of type <TYPE> and returns (bool, error)
//  2. slice of type []*<TYPE>
//
// Returns:
//  new filtered list and error
func Filter<FTYPE>PtrErr(f func(*<TYPE>) (bool, error), list []*<TYPE>) ([]*<TYPE>, error) {
	if f == nil {
		return []*<TYPE>{}, nil
	}
	var newList []*<TYPE>
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}
`
}

// FilterErr filters list based on function passed as 1st argument
func FilterErr() string {
	return `
// Filter<FTYPE>Err takes two arguments
//  1. Function: takes 1 argument of type <TYPE> and returns (bool, error)
//  2. slice of type []<TYPE>
//
// Returns:
//  new filtered list and error
func Filter<FTYPE>Err(f func(<TYPE>) (bool, error), list []<TYPE>) ([]<TYPE>, error) {
	if f == nil {
		return []<TYPE>{}, nil
	}
	var newList []<TYPE>
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}
`
}
