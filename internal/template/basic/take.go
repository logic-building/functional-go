package basic

// Take template to return n items in the list
func Take() string {
	return `
// Take<FTYPE> returns n items in the list
func Take<FTYPE>(n int, list []<TYPE>) []<TYPE> {
	if n < 0 {
		return []<TYPE>{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]<TYPE>, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}
`
}

// TakePtr template to return n items in the list
func TakePtr() string {
	return `
// Take<FTYPE>Ptr returns n items in the list
func Take<FTYPE>Ptr(n int, list []*<TYPE>) []*<TYPE> {
	if n < 0 {
		return []*<TYPE>{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*<TYPE>, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}
`
}
