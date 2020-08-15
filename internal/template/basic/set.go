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
`
}

// SetPtr is template.
func SetPtr() string {
	return `
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

// Set2 is template. template2 uses reflect.DeepEqual for struct which contains other than basic types such as list as member
func Set2() string {
	return `
// Set<FTYPE> returns a set of the distinct elements of coll.
func Set<FTYPE>(list []<TYPE>) []<TYPE> {
	if len(list) == 0 {
		return []<TYPE>{}
	}

	newList := []<TYPE>{list[0]}

	for i := 1; i < len(list); i++ {
		found := false
		for j := 0; j < len(newList); j++ {
			if reflect.DeepEqual(list[i], newList[j]) {
				found = true
				break
			}
		}
		if !found {
			newList = append(newList, list[i])
		}
	}
	return newList
}
`
}

// SetPtr2 is template. template2 uses reflect.DeepEqual for struct which contains other than basic types such as list as member
func SetPtr2() string {
	return `
// Set<FTYPE>Ptr returns a set of the distinct elements of coll.
func Set<FTYPE>Ptr(list []*<TYPE>) []*<TYPE> {
	if len(list) == 0 {
		return []*<TYPE>{}
	}

	newList := []*<TYPE>{list[0]}

	for i := 1; i < len(list); i++ {
		found := false
		for j := 0; j < len(newList); j++ {
			if reflect.DeepEqual(*list[i], *newList[j]) {
				found = true
				break
			}
		}
		if !found {
			newList = append(newList, list[i])
		}
	}
	return newList
}
`
}
