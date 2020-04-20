package basic

// FilterMapIO is template to generate itself for different combination of data type.
func FilterMapIO() string {
	return `
// FilterMap<FINPUT_TYPE><FOUTPUT_TYPE> filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - <INPUT_TYPE> and returns true/false.
//	2. Function: takes <INPUT_TYPE> as argument and returns <OUTPUT_TYPE>
// 	3. List of type <INPUT_TYPE>
//
// Returns:
//	New List of type <OUTPUT_TYPE>
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>(fFilter func(<INPUT_TYPE>) bool, fMap func(<INPUT_TYPE>) <OUTPUT_TYPE>, list []<INPUT_TYPE>) []<OUTPUT_TYPE> {
	if fFilter == nil || fMap == nil {
		return []<OUTPUT_TYPE>{}
	}
	var newList []<OUTPUT_TYPE>
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}
`
}

// FilterMapIOErr filters list based on function passed as 1st argument
func FilterMapIOErr() string {
	return `
// FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - <INPUT_TYPE> and returns (bool, error).
//	2. Function: takes <INPUT_TYPE> as argument and returns (<OUTPUT_TYPE>, error)
// 	3. List of type <INPUT_TYPE>
//
// Returns:
//	New List of type <OUTPUT_TYPE>, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(fFilter func(<INPUT_TYPE>) (bool, error), fMap func(<INPUT_TYPE>) (<OUTPUT_TYPE>, error), list []<INPUT_TYPE>) ([]<OUTPUT_TYPE>, error) {
	if fFilter == nil || fMap == nil {
		return []<OUTPUT_TYPE>{}, nil
	}
	var newList []<OUTPUT_TYPE>
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
