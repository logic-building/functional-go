package basic

// Dedupe is template to generate itself for different combination of data type.
func Dedupe() string {
	return `
// Dedupe<FTYPE> Returns a new list removing consecutive duplicates in list.
func Dedupe<FTYPE>(list []<TYPE>) []<TYPE> {
	var newList []<TYPE>

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// Dedupe<FTYPE>Ptr Returns a new list removing consecutive duplicates in list.
func Dedupe<FTYPE>Ptr(list []*<TYPE>) []*<TYPE> {
	var newList []*<TYPE>

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}
`
}
