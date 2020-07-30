package basic

// Equal is template to generate itself for different combination of data type.
func Equal() string {
	return `
// Equal<FTYPE>sP Returns true if both list are equal else returns false
func Equal<FTYPE>sP(list1, list2 []<TYPE>) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// Equal<FTYPE>sPPtr Returns true if both list are equal else returns false
func Equal<FTYPE>sPPtr(list1, list2 []*<TYPE>) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}
`
}
