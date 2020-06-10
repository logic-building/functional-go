package basic

// MapPtr applies the function(1st argument) on each item of the list and returns new list
func MapPtr() string {
	return `
// Map<FTYPE>Ptr takes 2 arguments.
//  1. A function input argument: <TYPE> and return type (<TYPE>)
//  2. A list of type []*<TYPE>
//
// Returns:
// 	([]*<TYPE>)
func Map<FTYPE>Ptr(f func(*<TYPE>) *<TYPE>, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	newList := make([]*<TYPE>, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
`
}

// MapPtrErr applies the function(1st argument) on each item of the list and returns new list
func MapPtrErr() string {
	return `
// Map<FTYPE>PtrErr takes 2 arguments:
//  1. A function input argument: *<TYPE> and return types (*<TYPE>, error)
//  2. A list of type []*<TYPE>
//
// Returns:
// 	([]*<TYPE>, error)
func Map<FTYPE>PtrErr(f func(*<TYPE>) (*<TYPE>, error), list []*<TYPE>) ([]*<TYPE>, error) {
	if f == nil {
		return []*<TYPE>{}, nil
	}
	newList := make([]*<TYPE>, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}
`
}

// MapErr applies the function(1st argument) on each item of the list and returns new list
func MapErr() string {
	return `
// Map<FTYPE>Err takes 2 arguments:
//  1. A function input argument: <TYPE> and return types (<TYPE>, error)
//  2. A list of type []<TYPE>
//
// Returns:
// 	([]<TYPE>, error)
func Map<FTYPE>Err(f func(<TYPE>) (<TYPE>, error), list []<TYPE>) ([]<TYPE>, error) {
	if f == nil {
		return []<TYPE>{}, nil
	}
	newList := make([]<TYPE>, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}
`
}
