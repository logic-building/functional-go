package basic

// Subset is template.
func Subset() string {
	return `
// Subset<FTYPE> returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func Subset<FTYPE>(list1, list2 []<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<TYPE>]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// Subset<FTYPE>Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func Subset<FTYPE>Ptr(list1, list2 []*<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<TYPE>]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}
`
}
