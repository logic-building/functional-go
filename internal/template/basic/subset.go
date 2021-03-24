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

	resultMap := make(map[<TYPE>]struct{})
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = struct{}{}
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
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

// SubsetPtr is template.
func SubsetPtr() string {
	return `
// Subset<FTYPE>Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func Subset<FTYPE>Ptr(list1, list2 []*<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<TYPE>]struct{})
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = struct{}{}
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
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

// Subset2 is template. template2 uses reflect.DeepEqual for struct which contains other than basic types such as list as member
func Subset2() string {
	return `
// Subset<FTYPE> returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored, order does not matter
func Subset<FTYPE>(list1, list2 []<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	for i := 0; i < len(list1); i++ {
		found := false
		for j := 0; j < len(list2); j++ {
			if reflect.DeepEqual(list1[i], list2[j]) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
`
}

// SubsetPtr2 is template. template2 uses reflect.DeepEqual for struct which contains other than basic types such as list as member
func SubsetPtr2() string {
	return `
// Subset<FTYPE>Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored, order does not matter
func Subset<FTYPE>Ptr(list1, list2 []*<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	for i := 0; i < len(list1); i++ {
		found := false
		for j := 0; j < len(list2); j++ {
			if reflect.DeepEqual(*list1[i], *list2[j]) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
`
}
