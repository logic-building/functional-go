package template

func Remove() string {
	return `
func Remove<CONDITIONAL_TYPE>(f func(<TYPE>) bool, list []<TYPE>) []<TYPE> {
	if f == nil {
		return []<TYPE>{}
	}
	var newList []<TYPE>
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}`
}
