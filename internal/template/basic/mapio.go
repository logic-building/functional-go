package basic

// Map is template to generate function(Map) for user defined data type
func MapIO() string {
	return `
func Map<FINPUT_TYPE><FOUTPUT_TYPE>(f func(<INPUT_TYPE>) <OUTPUT_TYPE>, list []<INPUT_TYPE>) []<OUTPUT_TYPE> {
	if f == nil {
		return []<OUTPUT_TYPE>{}
	}
	newList := make([]<OUTPUT_TYPE>, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}`
}
