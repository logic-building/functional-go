package basic

// Union<FTYPE> is template.
func Union() string {
	return `
// Union<FTYPE> return a set that is the union of the input sets
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

// Union<FTYPE>Ptr return a set that is the union of the input sets
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
