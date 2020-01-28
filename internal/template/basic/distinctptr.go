package basic

// DistinctInt removes duplicates.
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
