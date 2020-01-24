package basic

// Map<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func MapPtr() string {
	return `
func Map<FTYPE>Ptr(f func(*<TYPE>) *<TYPE>, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	newList := make([]*<TYPE>, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
`
}
