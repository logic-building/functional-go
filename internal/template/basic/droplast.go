package basic

// DropLast is template to generate itself for different combination of data type.
func DropLast() string {
	return `
// DropLast<FTYPE> drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLast<FTYPE>(list []<TYPE>) []<TYPE> {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []<TYPE>{}
	}

	newList := make([]<TYPE>, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}
`
}

// DropLastPtr is template to generate itself for different combination of data type.
func DropLastPtr() string {
	return `
// DropLast<FTYPE>Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLast<FTYPE>Ptr(list []*<TYPE>) []*<TYPE> {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*<TYPE>{}
	}

	newList := make([]*<TYPE>, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}
`
}
