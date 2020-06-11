package basic

// EveryPtr is template
func EveryPtr() string {
	return `
// Every<FTYPE>Ptr returns true if supplied function returns logical true for every item in the list
func Every<FTYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
`
}

// EveryPtrErr is template
func EveryPtrErr() string {
	return `
// Every<FTYPE>PtrErr returns true if supplied function returns logical true for every item in the list
func Every<FTYPE>PtrErr(f func(*<TYPE>) (bool, error), list []*<TYPE>) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}
`
}

// EveryErr is template
func EveryErr() string {
	return `
// Every<FTYPE>Err returns true if supplied function returns logical true for every item in the list
func Every<FTYPE>Err(f func(<TYPE>) (bool, error), list []<TYPE>) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}
`
}
