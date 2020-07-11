package basic

// Intersection is template.
func Intersection() string {
	return `
// Intersection<FTYPE> return a set that is the intersection of the input sets
// repeated value within list parameter will be ignored
func Intersection<FTYPE>(arrList ...[]<TYPE>) []<TYPE> {
	if arrList == nil {
		return []<TYPE>{}
	}

	resultMap := make(map[<TYPE>]bool)
	if len(arrList) == 1 {
		var newList []<TYPE>
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
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
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// Intersection<FTYPE>Ptr return a set that is the intersection of the input sets
// repeated value within list parameter will be ignored
func Intersection<FTYPE>Ptr(arrList ...[]*<TYPE>) []*<TYPE> {
	if arrList == nil {
		return []*<TYPE>{}
	}

	resultMap := make(map[<TYPE>]bool)
	if len(arrList) == 1 {
		var newList []*<TYPE>
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
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
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}
`
}
