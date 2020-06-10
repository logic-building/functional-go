package basic

// FilterMapErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List, error.
//  Empty list if all there parameters are nil or either of parameter is nil
//
func FilterMapErr() string {
	return `
// FilterMap<FTYPE>Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(<TYPE>) and returns (bool, error).
//	2. Function: takes <TYPE> as argument and returns (<TYPE>, error)
// 	3. Slice of type []<TYPE>
//
// Returns:
//	New List ([]<TYPE>, error).
//  Empty list if all there parameters are nil or either of parameter is nil
//
func FilterMap<FTYPE>Err(fFilter func(<TYPE>) (bool, error), fMap func(<TYPE>) (<TYPE>, error), list []<TYPE>) ([]<TYPE>, error) {
	if fFilter == nil || fMap == nil {
		return []<TYPE>{}, nil
	}
	var newList []<TYPE>
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
