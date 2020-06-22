package template

// SetStruct is template
func SetStruct() string {
	return `
// Union<FSTRUCT_NAME>By<FFIELD_NAME> return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func Union<FSTRUCT_NAME>By<FFIELD_NAME>(arrList ...[]<STRUCT_NAME>) []<STRUCT_NAME> {
	resultMap := make(map[<FIELD_TYPE>]bool)
	var resultArr []<STRUCT_NAME>
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[v.<FIELD_NAME>]
			if !ok {
				resultMap[v.<FIELD_NAME>] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// Union<FTYPE>Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func Union<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(arrList ...[]*<STRUCT_NAME>) []*<STRUCT_NAME> {
	resultMap := make(map[<FIELD_TYPE>]bool)
	var resultArr []*<STRUCT_NAME>
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[v.<FIELD_NAME>]
			if !ok {
				resultMap[v.<FIELD_NAME>] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// Intersection<FSTRUCT_NAME>By<FFIELD_NAME> return a set that is the intersection of the input sets
// repeated value within list parameter will be ignored
func Intersection<FSTRUCT_NAME>By<FFIELD_NAME>(arrList ...[]<STRUCT_NAME>) []<STRUCT_NAME> {
	if arrList == nil {
		return []<STRUCT_NAME>{}
	}

	resultMap := make(map[<FIELD_TYPE>]bool)
	if len(arrList) == 1 {
		var newList []<STRUCT_NAME>
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i].<FIELD_NAME>]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i].<FIELD_NAME>] = true
			}
		}
		return newList
	}

	var newList []<STRUCT_NAME>
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i].<FIELD_NAME> == v.<FIELD_NAME> {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i].<FIELD_NAME>]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i].<FIELD_NAME>] = true
			}
		}
	}
	return newList
}

// Intersection<FSTRUCT_NAME>By<FFIELD_NAME>Ptr return a set that is the intersection of the input sets
// repeated value within list parameter will be ignored
func Intersection<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(arrList ...[]*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if arrList == nil {
		return []*<STRUCT_NAME>{}
	}

	resultMap := make(map[<FIELD_TYPE>]bool)
	if len(arrList) == 1 {
		var newList []*<STRUCT_NAME>
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i].<FIELD_NAME>]
			if !ok {
				resultMap[arrList[0][i].<FIELD_NAME>] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*<STRUCT_NAME>
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i].<FIELD_NAME> == v.<FIELD_NAME> {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i].<FIELD_NAME>]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i].<FIELD_NAME>] = true
			}
		}
	}
	return newList
}

// Difference<FSTRUCT_NAME>By<FFIELD_NAME> returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func Difference<FSTRUCT_NAME>By<FFIELD_NAME>(arrList ...[]<STRUCT_NAME>) []<STRUCT_NAME> {
	if arrList == nil {
		return []<STRUCT_NAME>{}
	}

	resultMap := make(map[<FIELD_TYPE>]bool)
	if len(arrList) == 1 {
		var newList []<STRUCT_NAME>
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i].<FIELD_NAME>]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i].<FIELD_NAME>] = true
			}
		}
		return newList
	}

	var newList []<STRUCT_NAME>
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i].<FIELD_NAME> == v.<FIELD_NAME> {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i].<FIELD_NAME>]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i].<FIELD_NAME>] = true
			}
		}
	}
	return newList
}

// Difference<FSTRUCT_NAME>By<FFIELD_NAME>Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func Difference<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(arrList ...[]*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if arrList == nil {
		return []*<STRUCT_NAME>{}
	}

	resultMap := make(map[<FIELD_TYPE>]bool)
	if len(arrList) == 1 {
		var newList []*<STRUCT_NAME>
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i].<FIELD_NAME>]
			if !ok {
				resultMap[arrList[0][i].<FIELD_NAME>] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*<STRUCT_NAME>
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i].<FIELD_NAME> == v.<FIELD_NAME> {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i].<FIELD_NAME>]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i].<FIELD_NAME>] = true
			}
		}
	}
	return newList
}

// Subset<FSTRUCT_NAME>By<FFIELD_NAME> returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func Subset<FSTRUCT_NAME>By<FFIELD_NAME>(list1, list2 []<STRUCT_NAME>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<FIELD_TYPE>]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i].<FIELD_NAME>]
		if !ok {
			found := false
			resultMap[list1[i].<FIELD_NAME>] = true
			for j := 0; j < len(list2); j++ {
				if list1[i].<FIELD_NAME> == list2[j].<FIELD_NAME> {
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

// Subset<FSTRUCT_NAME>By<FFIELD_NAME>Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func Subset<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(list1, list2 []*<STRUCT_NAME>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<FIELD_TYPE>]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i].<FIELD_NAME>]
		if !ok {
			found := false
			resultMap[list1[i].<FIELD_NAME>] = true
			for j := 0; j < len(list2); j++ {
				if list1[i].<FIELD_NAME> == list2[j].<FIELD_NAME> {
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

// Superset<FSTRUCT_NAME>By<FFIELD_NAME> returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func Superset<FSTRUCT_NAME>By<FFIELD_NAME>(list1, list2 []<STRUCT_NAME>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<FIELD_TYPE>]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i].<FIELD_NAME>]
		if !ok {
			found := false
			resultMap[list2[i].<FIELD_NAME>] = true
			for j := 0; j < len(list1); j++ {
				if list2[i].<FIELD_NAME> == list1[j].<FIELD_NAME> {
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

// Superset<FSTRUCT_NAME>By<FFIELD_NAME>Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func Superset<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(list1, list2 []*<STRUCT_NAME>) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[<FIELD_TYPE>]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i].<FIELD_NAME>]
		if !ok {
			found := false
			resultMap[list2[i].<FIELD_NAME>] = true
			for j := 0; j < len(list1); j++ {
				if list2[i].<FIELD_NAME> == list1[j].<FIELD_NAME> {
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

// Set<FSTRUCT_NAME>By<FFIELD_NAME> returns a set of the distinct elements of coll.
func Set<FSTRUCT_NAME>By<FFIELD_NAME>(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if list == nil || len(list) == 0 {
		return []<STRUCT_NAME>{}
	}

	resultMap := make(map[<FIELD_TYPE>]bool)
	newList := []<STRUCT_NAME>{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i].<FIELD_NAME>]
		if !ok {
			resultMap[list[i].<FIELD_NAME>] = true
			newList = append(newList, list[i])
		}
	}
	return newList
}

// Set<FSTRUCT_NAME>By<FFIELD_NAME>Ptr returns a set of the distinct elements of coll.
func Set<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if list == nil || len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}

	resultMap := make(map[<FIELD_TYPE>]bool)
	newList := []*<STRUCT_NAME>{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i].<FIELD_NAME>]
		if !ok {
			resultMap[list[i].<FIELD_NAME>] = true
			newList = append(newList, list[i])
		}
	}
	return newList
}`
}
