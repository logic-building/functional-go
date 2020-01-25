package basic

// FilterInt filters list based on function passed as 1st argument
func FilterPtr() string {
	return `
func Filter<FTYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
`
}
