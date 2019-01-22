package template

func Map() string {
	return `
func Map<CONDITIONAL_TYPE>(f func(<TYPE>) <TYPE>, list []<TYPE>) []<TYPE> {
	if f == nil {
		return []<TYPE>{}
	}
	newList := make([]<TYPE>, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}`
}
