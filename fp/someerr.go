package fp

// SomeIntErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeIntErr(f func(int) (bool, error), list []int) (bool, error) {
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

// SomeInt64Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeInt64Err(f func(int64) (bool, error), list []int64) (bool, error) {
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

// SomeInt32Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeInt32Err(f func(int32) (bool, error), list []int32) (bool, error) {
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

// SomeInt16Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeInt16Err(f func(int16) (bool, error), list []int16) (bool, error) {
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

// SomeInt8Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeInt8Err(f func(int8) (bool, error), list []int8) (bool, error) {
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

// SomeUintErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUintErr(f func(uint) (bool, error), list []uint) (bool, error) {
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

// SomeUint64Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUint64Err(f func(uint64) (bool, error), list []uint64) (bool, error) {
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

// SomeUint32Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUint32Err(f func(uint32) (bool, error), list []uint32) (bool, error) {
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

// SomeUint16Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUint16Err(f func(uint16) (bool, error), list []uint16) (bool, error) {
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

// SomeUint8Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeUint8Err(f func(uint8) (bool, error), list []uint8) (bool, error) {
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

// SomeStrErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeStrErr(f func(string) (bool, error), list []string) (bool, error) {
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

// SomeBoolErr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeBoolErr(f func(bool) (bool, error), list []bool) (bool, error) {
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

// SomeFloat32Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeFloat32Err(f func(float32) (bool, error), list []float32) (bool, error) {
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

// SomeFloat64Err finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	(bool,err).
//	True if condition satisfies, else false
func SomeFloat64Err(f func(float64) (bool, error), list []float64) (bool, error) {
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
