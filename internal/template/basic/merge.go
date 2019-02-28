package basic

// Merge is template to generate itself for different combination of data type.
func Merge() string {
	return `
// Merge<FINPUT_TYPE1><FINPUT_TYPE2> takes two inputs: map[<INPUT_TYPE1>]<INPUT_TYPE2> and map[<INPUT_TYPE1>]<INPUT_TYPE2> and merge two maps and returns a new map[<INPUT_TYPE1>]<INPUT_TYPE2>.
func Merge<FINPUT_TYPE1><FINPUT_TYPE2>(map1, map2 map[<INPUT_TYPE1>]<INPUT_TYPE2>) map[<INPUT_TYPE1>]<INPUT_TYPE2> {
	if map1 == nil && map2 == nil {
		return map[<INPUT_TYPE1>]<INPUT_TYPE2>{}
	}

	newMap := make(map[<INPUT_TYPE1>]<INPUT_TYPE2>)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}
`
}
