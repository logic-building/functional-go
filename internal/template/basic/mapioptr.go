package basic

// MapIO is template to generate function(Map) for user defined data type
func MapIOPtr() string {
	return `
// Map<FINPUT_TYPE><FOUTPUT_TYPE> takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func Map<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(f func(*<INPUT_TYPE>) *<OUTPUT_TYPE>, list []*<INPUT_TYPE>) []*<OUTPUT_TYPE> {
	if f == nil {
		return []*<OUTPUT_TYPE>{}
	}
	newList := make([]*<OUTPUT_TYPE>, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
`
}
