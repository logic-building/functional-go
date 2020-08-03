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

// EqualMapType1Type2 is template to generate itself for different combination of data type.
func EqualMapType1Type2() string {
	return `
// EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P Returns true if both maps are equal else returns false
func EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>P(map1, map2 map[<INPUT_TYPE1>]<INPUT_TYPE2>) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr Returns true if both maps are equal else returns false
func EqualMap<FINPUT_TYPE1><FINPUT_TYPE2>PPtr(map1, map2 map[*<INPUT_TYPE1>]*<INPUT_TYPE2>) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
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
