package basic

// Distinct is template.
func Distinct() string {
	return `
// Distinct<FTYPE> removes duplicates.
func Distinct<FTYPE>(list []<TYPE>) []<TYPE> {
	var newList []<TYPE>
	s := make(map[<TYPE>]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}
`
}

// Distinct2 is template for struct which holds complex member such as list.
func Distinct2() string {
	return `
// Distinct<FTYPE> removes duplicates.
func Distinct<FTYPE>(list []<TYPE>) []<TYPE> {
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

// DistinctPtr is template.
func DistinctPtr() string {
	return `
// Distinct<FTYPE>Ptr removes duplicates.
func Distinct<FTYPE>Ptr(list []*<TYPE>) []*<TYPE> {
	var newList []*<TYPE>
	s := make(map[<TYPE>]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}
`
}

// DistinctPtr2 is template for struct which holds complex member such as list.
func DistinctPtr2() string {
	return `
func Distinct<FTYPE>Ptr(list []*<TYPE>) []*<TYPE> {
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

// Distinct check
// DistinctP is template.
func DistinctP() string {
	return `
// Distinct<FTYPE>P returns true if no two of the arguments are =
func Distinct<FTYPE>P(list []<TYPE>) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[<TYPE>]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}
`
}

// DistinctP2 is template.
func DistinctP2() string {
	return `
// Distinct<FTYPE>P returns true if no two of the arguments are =
func Distinct<FTYPE>P(list []<TYPE>) bool {
	if len(list) == 0 {
		return false
	}

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if reflect.DeepEqual(list[i], list[j]) {
				return false
			}
		}
	}
	return true
}
`
}

// DistinctPPtr is template.
func DistinctPPtr() string {
	return `
// Distinct<FTYPE>PPtr returns true if no two of the arguments are =
func Distinct<FTYPE>PPtr(list []*<TYPE>) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[<TYPE>]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}
`
}

// DistinctPPtr2 is template.
func DistinctPPtr2() string {
	return `
// Distinct<FTYPE>PPtr returns true if no two of the arguments are =
func Distinct<FTYPE>PPtr(list []*<TYPE>) bool {
	if len(list) == 0 {
		return false
	}

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if reflect.DeepEqual(*list[i], *list[j]) {
				return false
			}
		}
	}
	return true
}
`
}
