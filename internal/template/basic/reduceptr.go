package basic

// ReducePtr reduces a list to a single value by combining elements via a supplied function
func ReducePtr() string {
	return `
// Reduce<FTYPE>Ptr reduces a list to a single value by combining elements via a supplied function
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

// ReducePtrErr reduces a list to a single value by combining elements via a supplied function
func ReducePtrErr() string {
	return `
// Reduce<FTYPE>PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*<TYPE>, *<TYPE>)
//	B. list of type []*<TYPE>
// 	C. initializer (optional of type <TYPE>)
//
// Returns:
//	single value, error: (*<TYPE>, error)
func Reduce<FTYPE>PtrErr(f func(*<TYPE>, *<TYPE>) (*<TYPE>, error), list []*<TYPE>, initializer ...<TYPE>) (*<TYPE>, error) {
	var initVal <TYPE>
	var init *<TYPE> = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return Reduce<FTYPE>PtrErr(f, list[1:], *r)
}
`
}

// ReduceErr reduces a list to a single value by combining elements via a supplied function
func ReduceErr() string {
	return `
// Reduce<FTYPE>Err reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (<TYPE>, <TYPE>)
//	B. list of type []<TYPE>
// 	C. initializer (optional of type <TYPE>)
//
// Returns:
//	single value, error: (<TYPE>, error)
func Reduce<FTYPE>Err(f func(<TYPE>, <TYPE>) (<TYPE>, error), list []<TYPE>, initializer ...<TYPE>) (<TYPE>, error) {
	var initVal <TYPE>
	var init <TYPE> = initVal
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return r, err
	}
	return Reduce<FTYPE>Err(f, list[1:], r)
}
`
}
