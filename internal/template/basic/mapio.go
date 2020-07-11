package basic

// MapIO is template to generate function(Map) for user defined data type
func MapIO() string {
	return `
// Map<FINPUT_TYPE><FOUTPUT_TYPE> takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func Map<FINPUT_TYPE><FOUTPUT_TYPE>(f func(<INPUT_TYPE>) <OUTPUT_TYPE>, list []<INPUT_TYPE>) []<OUTPUT_TYPE> {
	if f == nil {
		return []<OUTPUT_TYPE>{}
	}
	newList := make([]<OUTPUT_TYPE>, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
`
}

// MapIOErr is template to generate function(Map) for user defined data type
func MapIOErr() string {
	return `
// Map<FINPUT_TYPE><FOUTPUT_TYPE>Err takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func Map<FINPUT_TYPE><FOUTPUT_TYPE>Err(f func(<INPUT_TYPE>) (<OUTPUT_TYPE>, error), list []<INPUT_TYPE>) ([]<OUTPUT_TYPE>, error) {
	if f == nil {
		return []<OUTPUT_TYPE>{}, nil
	}
	newList := make([]<OUTPUT_TYPE>, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}
`
}
