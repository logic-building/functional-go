package template

func TakeWhile() string {
	return `
func TakeWhile<CONDITIONAL_TYPE>(f func(<TYPE>) bool, list []<TYPE>) []<TYPE> {
	if f == nil {
		return []<TYPE>{}
	}
	var newList []<TYPE>
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}`
}
