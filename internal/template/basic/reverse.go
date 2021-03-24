package basic

// Reverse reverse the list.
func Reverse() string {
	return `
// Reverse<FTYPE>s reverse the list
func Reverse<FTYPE>s(list []<TYPE>) []<TYPE> {
	newList := make([]<TYPE>, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}
`
}

// ReversePtr reverse the list.
func ReversePtr() string {
	return `
// Reverse<FTYPE>sPtr reverse the list
func Reverse<FTYPE>sPtr(list []*<TYPE>) []*<TYPE> {
	newList := make([]*<TYPE>, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}
`
}
