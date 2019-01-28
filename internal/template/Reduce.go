package template

func Reduce() string {
	return `
func Reduce<CONDITIONAL_TYPE>(f func(<TYPE>, <TYPE>) <TYPE>, list []<TYPE>, initializer ...<TYPE>) <TYPE> {
	var init <TYPE> 
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}
	
	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return Reduce<CONDITIONAL_TYPE>(f, list[1:], r)
}`
}
