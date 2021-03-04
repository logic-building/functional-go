package basic

// Difference is template.
func Difference() string {
	return `
// Difference<FTYPE> returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func Difference<FTYPE>(arrList ...[]<TYPE>) []<TYPE> {
	if arrList == nil {
		return []<TYPE>{}
	}

	resultMap := make(map[<TYPE>]struct{})
	if len(arrList) == 1 {
		var newList []<TYPE>
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []<TYPE>
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}
`
}

// DifferencePtr is template.
func DifferencePtr() string {
	return `
// Difference<FTYPE>Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func Difference<FTYPE>Ptr(arrList ...[]*<TYPE>) []*<TYPE> {
	if arrList == nil {
		return []*<TYPE>{}
	}

	resultMap := make(map[<TYPE>]struct{})
	if len(arrList) == 1 {
		var newList []*<TYPE>
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*<TYPE>
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}
`
}

// Difference2 is template. template2 uses reflect.DeepEqual for struct which contains other than basic types such as list as member
func Difference2() string {
	return `
// Difference<FTYPE> returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func Difference<TYPE>(arrList ...[]<TYPE>) []<TYPE> {
	if arrList == nil {
		return []<TYPE>{}
	}

	var newList []<TYPE>

	if len(arrList) == 1 {
		for i := 0; i < len(arrList[0]); i++ {
			found := false
			for j := 0; j < len(newList); j++ {
				if reflect.DeepEqual(newList[j], arrList[0][i]) {
					found = true
					break
				}
			}
			if !found {
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if reflect.DeepEqual(arrList[0][i], v) {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			found := false
			for j := 0; j < len(newList); j++ {
				if reflect.DeepEqual(newList[j], arrList[0][i]) {
					found = true
					break
				}
			}
			if !found {
				newList = append(newList, arrList[0][i])
			}
		}
	}
	return newList
}
`
}

// DifferencePtr2 is template. template2 uses reflect.DeepEqual for struct which contains other than basic types such as list as member
func DifferencePtr2() string {
	return `
// Difference<FTYPE>Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func Difference<FTYPE>Ptr(arrList ...[]*<TYPE>) []*<TYPE> {
	if arrList == nil {
		return []*<TYPE>{}
	}

	var newList []*<TYPE>

	if len(arrList) == 1 {
		for i := 0; i < len(arrList[0]); i++ {
			found := false
			for j := 0; j < len(newList); j++ {
				if reflect.DeepEqual(*newList[j], *arrList[0][i]) {
					found = true
					break
				}
			}
			if !found {
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if reflect.DeepEqual(*arrList[0][i], *v) {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			found := false
			for j := 0; j < len(newList); j++ {
				if reflect.DeepEqual(*newList[j], *arrList[0][i]) {
					found = true
					break
				}
			}
			if !found {
				newList = append(newList, arrList[0][i])
			}
		}
	}
	return newList
}
`
}
