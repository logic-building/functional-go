package basic

// SomePtr is template to generate function(Some) for user defined data type
func SomePtr() string {
	return `
// Some<FTYPE>Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func Some<FTYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
`
}

// SomePtrErr is template to generate function(Some) for user defined data type
func SomePtrErr() string {
	return `
// Some<FTYPE>PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func Some<FTYPE>PtrErr(f func(*<TYPE>) (bool, error), list []*<TYPE>) (bool, error) {
	if f == nil {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if r {
			return true, nil
		}
	}
	return false, nil
}
`
}

// SomeErr is template to generate function(Some) for user defined data type
func SomeErr() string {
	return `
// Some<FTYPE>Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func Some<FTYPE>Err(f func(<TYPE>) (bool, error), list []<TYPE>) (bool, error) {
	if f == nil {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if r {
			return true, nil
		}
	}
	return false, nil
}
`
}
