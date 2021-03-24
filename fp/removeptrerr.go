package fp

// RemoveIntPtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *int and return types(bool, error)
//	2. List of type: []*int
//
// Returns:
//	New list and error: ([]*int, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveIntPtrErr(f func(*int) (bool, error), list []*int) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveInt64PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *int64 and return types(bool, error)
//	2. List of type: []*int64
//
// Returns:
//	New list and error: ([]*int64, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveInt64PtrErr(f func(*int64) (bool, error), list []*int64) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveInt32PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *int32 and return types(bool, error)
//	2. List of type: []*int32
//
// Returns:
//	New list and error: ([]*int32, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveInt32PtrErr(f func(*int32) (bool, error), list []*int32) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveInt16PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *int16 and return types(bool, error)
//	2. List of type: []*int16
//
// Returns:
//	New list and error: ([]*int16, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveInt16PtrErr(f func(*int16) (bool, error), list []*int16) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveInt8PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *int8 and return types(bool, error)
//	2. List of type: []*int8
//
// Returns:
//	New list and error: ([]*int8, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveInt8PtrErr(f func(*int8) (bool, error), list []*int8) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveUintPtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *uint and return types(bool, error)
//	2. List of type: []*uint
//
// Returns:
//	New list and error: ([]*uint, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUintPtrErr(f func(*uint) (bool, error), list []*uint) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveUint64PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *uint64 and return types(bool, error)
//	2. List of type: []*uint64
//
// Returns:
//	New list and error: ([]*uint64, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUint64PtrErr(f func(*uint64) (bool, error), list []*uint64) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveUint32PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *uint32 and return types(bool, error)
//	2. List of type: []*uint32
//
// Returns:
//	New list and error: ([]*uint32, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUint32PtrErr(f func(*uint32) (bool, error), list []*uint32) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveUint16PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *uint16 and return types(bool, error)
//	2. List of type: []*uint16
//
// Returns:
//	New list and error: ([]*uint16, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUint16PtrErr(f func(*uint16) (bool, error), list []*uint16) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveUint8PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *uint8 and return types(bool, error)
//	2. List of type: []*uint8
//
// Returns:
//	New list and error: ([]*uint8, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUint8PtrErr(f func(*uint8) (bool, error), list []*uint8) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveStrPtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *string and return types(bool, error)
//	2. List of type: []*string
//
// Returns:
//	New list and error: ([]*string, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveStrPtrErr(f func(*string) (bool, error), list []*string) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveBoolPtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *bool and return types(bool, error)
//	2. List of type: []*bool
//
// Returns:
//	New list and error: ([]*bool, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveBoolPtrErr(f func(*bool) (bool, error), list []*bool) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveFloat32PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *float32 and return types(bool, error)
//	2. List of type: []*float32
//
// Returns:
//	New list and error: ([]*float32, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveFloat32PtrErr(f func(*float32) (bool, error), list []*float32) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// RemoveFloat64PtrErr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function: input type *float64 and return types(bool, error)
//	2. List of type: []*float64
//
// Returns:
//	New list and error: ([]*float64, error)
//	Empty list if both of arguments are nil or either one is nil.
func RemoveFloat64PtrErr(f func(*float64) (bool, error), list []*float64) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}
