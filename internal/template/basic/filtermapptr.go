package basic

// FilterMapPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapPtr() string {
	return `
// FilterMap<FTYPE>Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes <TYPE> as argument and returns (<TYPE>)
// 	3. Slice of type []*<TYPE>
//
// Returns:
//	New List of type []*<TYPE>.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMap<FTYPE>Ptr(fFilter func(*<TYPE>) bool, fMap func(*<TYPE>) *<TYPE>, list []*<TYPE>) []*<TYPE> {
	if fFilter == nil || fMap == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}
`
}

// FilterMapPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List, error.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapPtrErr() string {
	return `
// FilterMap<FTYPE>PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*<TYPE>) and returns (bool, error).
//	2. Function: takes *<TYPE> as argument and returns (*<TYPE>, error)
// 	3. Slice of type []*<TYPE>
//
// Returns:
//	New List ([]*<TYPE>, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMap<FTYPE>PtrErr(fFilter func(*<TYPE>) (bool, error), fMap func(*<TYPE>) (*<TYPE>, error), list []*<TYPE>) ([]*<TYPE>, error) {
	if fFilter == nil || fMap == nil {
		return []*<TYPE>{}, nil
	}
	var newList []*<TYPE>
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}
`
}
