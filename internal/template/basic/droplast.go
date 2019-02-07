package basic

func DropLast() string {
	return `
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
}`
}
