package basic

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
