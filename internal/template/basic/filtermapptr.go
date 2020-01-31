package basic

// FilterMapInt filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
func FilterMapPtr() string {
	return `
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
