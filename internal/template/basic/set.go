package basic

// Set is template.
func Set() string {
	return `
// Set<FTYPE> returns a set of the distinct elements of coll.
func Set<FTYPE>(list []<TYPE>) []<TYPE> {
	if list == nil || len(list) == 0 {
		return []<TYPE>{}
	}

	resultMap := make(map[<TYPE>]bool)
	newList := []<TYPE>{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = true
			newList = append(newList, list[i])
		}
	}
	return newList
}

// Set<FTYPE>Ptr returns a set of the distinct elements of coll.
func Set<FTYPE>Ptr(list []*<TYPE>) []*<TYPE> {
	if list == nil || len(list) == 0 {
		return []*<TYPE>{}
	}

	resultMap := make(map[<TYPE>]bool)
	newList := []*<TYPE>{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = true
			newList = append(newList, list[i])
		}
	}
	return newList
}
`
}
