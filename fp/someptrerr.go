package fp

// SomeIntPtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeIntPtrErr(f func(*int) (bool, error), list []*int) (bool, error) {
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

// SomeInt64PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeInt64PtrErr(f func(*int64) (bool, error), list []*int64) (bool, error) {
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

// SomeInt32PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeInt32PtrErr(f func(*int32) (bool, error), list []*int32) (bool, error) {
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

// SomeInt16PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeInt16PtrErr(f func(*int16) (bool, error), list []*int16) (bool, error) {
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

// SomeInt8PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeInt8PtrErr(f func(*int8) (bool, error), list []*int8) (bool, error) {
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

// SomeUintPtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUintPtrErr(f func(*uint) (bool, error), list []*uint) (bool, error) {
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

// SomeUint64PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUint64PtrErr(f func(*uint64) (bool, error), list []*uint64) (bool, error) {
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

// SomeUint32PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUint32PtrErr(f func(*uint32) (bool, error), list []*uint32) (bool, error) {
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

// SomeUint16PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUint16PtrErr(f func(*uint16) (bool, error), list []*uint16) (bool, error) {
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

// SomeUint8PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUint8PtrErr(f func(*uint8) (bool, error), list []*uint8) (bool, error) {
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

// SomeStrPtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeStrPtrErr(f func(*string) (bool, error), list []*string) (bool, error) {
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

// SomeBoolPtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeBoolPtrErr(f func(*bool) (bool, error), list []*bool) (bool, error) {
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

// SomeFloat32PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeFloat32PtrErr(f func(*float32) (bool, error), list []*float32) (bool, error) {
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

// SomeFloat64PtrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeFloat64PtrErr(f func(*float64) (bool, error), list []*float64) (bool, error) {
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
