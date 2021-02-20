package basic

// Keys is template to generate itself for different combination of data type.
func Keys() string {
	return `
// Keys<FINPUT_TYPE1><FINPUT_TYPE2> returns a slice of map's keys
func Keys<FINPUT_TYPE1><FINPUT_TYPE2>(m map[<INPUT_TYPE1>]<INPUT_TYPE2>) []<INPUT_TYPE1> {
	keys := make([]<INPUT_TYPE1>, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
`
}
