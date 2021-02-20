package basic

// Union is template.
func Union() string {
	return `
// Union<FTYPE> return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func Union<FTYPE>(arrList ...[]<TYPE>) []<TYPE> {
	resultMap := make(map[<TYPE>]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]<TYPE>, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}
`
}

// UnionPtr is template.
func UnionPtr() string {
	return `
// Union<FTYPE>Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func Union<FTYPE>Ptr(arrList ...[]*<TYPE>) []*<TYPE> {
	resultMap := make(map[<TYPE>]bool)
	var resultArr []*<TYPE>
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}
`
}

// For struct. template2 uses reflect.DeepEqual for struct which contains other than basic types such as list as member

// Union2 is template.
func Union2() string {
	return `
// Union<FTYPE> return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func Union<FTYPE>(arrList ...[]<TYPE>) []<TYPE> {
	var newList []<TYPE>

	for _, arr := range arrList {
		for _, v := range arr {
			found := false
			for i := 0; i < len(newList); i++ {
				if reflect.DeepEqual(newList[i], v) {
					found = true
					break
				}
			}
			if !found {
				newList = append(newList, v)
			}
		}
	}
	return newList
}
`
}

// Union2Ptr is template.
func Union2Ptr() string {
	return `
// Union<FTYPE>Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func Union<FTYPE>Ptr(arrList ...[]*<TYPE>) []*<TYPE> {
	var newList []*<TYPE>

	for _, arr := range arrList {
		for _, v := range arr {
			found := false
			for i := 0; i < len(newList); i++ {
				if reflect.DeepEqual(*newList[i], *v) {
					found = true
					break
				}
			}
			if !found {
				newList = append(newList, v)
			}
		}
	}
	return newList
}
`
}
