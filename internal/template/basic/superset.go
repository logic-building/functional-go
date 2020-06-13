package basic

// Superset is template.
func Superset() string {
	return `
// Subset<FTYPE> returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func Superset<FTYPE>(list1, list2 []<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<TYPE>]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// Superset<FTYPE>Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func Superset<FTYPE>Ptr(list1, list2 []*<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<TYPE>]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}
`
}
