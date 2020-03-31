package template

// FilterMap is template to generate function(FilterMap) for user defined data type
func FilterMap() string {
	return `
func FilterMap<CONDITIONAL_TYPE>(fFilter func(<TYPE>) bool, fMap func(<TYPE>) <TYPE>, list []<TYPE>) []<TYPE> {
	if fFilter == nil || fMap == nil {
		return []<TYPE>{}
	}
	var newList []<TYPE>
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}
`
}

// FilterMapPtr is template to generate function(FilterMap) for user defined data type
func FilterMapPtr() string {
	return `
func FilterMap<CONDITIONAL_TYPE>Ptr(fFilter func(*<TYPE>) bool, fMap func(*<TYPE>) *<TYPE>, list []*<TYPE>) []*<TYPE> {
	if fFilter == nil || fMap == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}
`
}
