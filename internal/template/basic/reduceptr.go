package basic

// Reduce<FTYPE> reduces a list to a single value by combining elements via a supplied function
func ReducePtr() string {
	return `
// Reduce<FTYPE> reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.

func Reduce<FTYPE>Ptr(f func(*<TYPE>, *<TYPE>) *<TYPE>, list []*<TYPE>, initializer ...<TYPE>) *<TYPE> {
	var initVal <TYPE>
	var init *<TYPE> = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return Reduce<FTYPE>Ptr(f, list[1:], *r)
}
`
}
