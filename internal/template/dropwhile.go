package template

// DropWhile is template to generate function(DropWhhile) for user defined data type
func DropWhile() string {
	return `
func DropWhile<CONDITIONAL_TYPE>(f func(<TYPE>) bool, list []<TYPE>) []<TYPE> {
	if f == nil {
		return []<TYPE>{}
	}
	var newList []<TYPE>
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]<TYPE>, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}
`
}

// DropWhilePtr is template to generate function(DropWhhile) for user defined data type
func DropWhilePtr() string {
	return `
func DropWhile<CONDITIONAL_TYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}
	var newList []*<TYPE>
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*<TYPE>, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}
`
}
