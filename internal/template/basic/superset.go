package basic

// Superset is template.
func Superset() string {
	return `
// Superset<FTYPE> returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func Superset<FTYPE>(list1, list2 []<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<TYPE>]struct{})

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = struct{}{}
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

// SupersetPtr is template.
func SupersetPtr() string {
	return `
// Superset<FTYPE>Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func Superset<FTYPE>Ptr(list1, list2 []*<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<TYPE>]struct{})

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = struct{}{}
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

// Superset2 is template. template2 uses reflect.DeepEqual for struct which contains other than basic types such as list as member
func Superset2() string {
	return `
// Superset<FTYPE> returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored, order does not matter
func Superset<FTYPE>(list1, list2 []<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	for i := 0; i < len(list2); i++ {
		found := false
		for j := 0; j < len(list1); j++ {
			if reflect.DeepEqual(list2[i], list1[j]) {
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

// SupersetPtr2 is template. template2 uses reflect.DeepEqual for struct which contains other than basic types such as list as member
func SupersetPtr2() string {
	return `
// Superset<FTYPE>Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored, order does not matter
func Superset<FTYPE>Ptr(list1, list2 []*<TYPE>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	for i := 0; i < len(list2); i++ {
		found := false
		for j := 0; j < len(list1); j++ {
			if reflect.DeepEqual(*list2[i], *list1[j]) {
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
