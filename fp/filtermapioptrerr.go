package fp

// FilterMapIntInt64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*int64, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt64PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*int64, error), list []*int) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntInt32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*int32, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt32PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*int32, error), list []*int) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntInt16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*int16, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt16PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*int16, error), list []*int) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntInt8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*int8, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt8PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*int8, error), list []*int) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntUintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*uint, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUintPtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*uint, error), list []*int) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntUint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*uint64, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint64PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*uint64, error), list []*int) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntUint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*uint32, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint32PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*uint32, error), list []*int) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntUint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*uint16, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint16PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*uint16, error), list []*int) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntUint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*uint8, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint8PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*uint8, error), list []*int) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntStrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*string, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntStrPtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*string, error), list []*int) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntBoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*bool, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntBoolPtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*bool, error), list []*int) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntFloat32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*float32, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntFloat32PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*float32, error), list []*int) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapIntFloat64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns (bool, error).
//	2. Function: takes *int as argument and returns (*float64, error)
// 	3. List of type *int
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntFloat64PtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*float64, error), list []*int) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*int, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64IntPtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*int, error), list []*int64) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64Int32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*int32, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int32PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*int32, error), list []*int64) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64Int16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*int16, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int16PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*int16, error), list []*int64) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64Int8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*int8, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int8PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*int8, error), list []*int64) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*uint, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64UintPtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*uint, error), list []*int64) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64Uint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*uint64, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint64PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*uint64, error), list []*int64) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64Uint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*uint32, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint32PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*uint32, error), list []*int64) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64Uint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*uint16, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint16PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*uint16, error), list []*int64) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64Uint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*uint8, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint8PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*uint8, error), list []*int64) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*string, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64StrPtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*string, error), list []*int64) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*bool, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64BoolPtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*bool, error), list []*int64) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64Float32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*float32, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Float32PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*float32, error), list []*int64) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt64Float64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*float64, error)
// 	3. List of type *int64
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Float64PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*float64, error), list []*int64) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*int, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32IntPtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*int, error), list []*int32) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32Int64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*int64, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int64PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*int64, error), list []*int32) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32Int16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*int16, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int16PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*int16, error), list []*int32) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32Int8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*int8, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int8PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*int8, error), list []*int32) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*uint, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32UintPtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*uint, error), list []*int32) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32Uint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*uint64, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint64PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*uint64, error), list []*int32) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32Uint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*uint32, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint32PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*uint32, error), list []*int32) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32Uint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*uint16, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint16PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*uint16, error), list []*int32) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32Uint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*uint8, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint8PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*uint8, error), list []*int32) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*string, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32StrPtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*string, error), list []*int32) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*bool, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32BoolPtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*bool, error), list []*int32) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32Float32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*float32, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Float32PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*float32, error), list []*int32) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt32Float64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*float64, error)
// 	3. List of type *int32
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Float64PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*float64, error), list []*int32) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*int, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16IntPtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*int, error), list []*int16) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16Int64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*int64, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int64PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*int64, error), list []*int16) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16Int32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*int32, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int32PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*int32, error), list []*int16) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16Int8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*int8, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int8PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*int8, error), list []*int16) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*uint, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16UintPtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*uint, error), list []*int16) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16Uint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*uint64, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint64PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*uint64, error), list []*int16) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16Uint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*uint32, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint32PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*uint32, error), list []*int16) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16Uint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*uint16, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint16PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*uint16, error), list []*int16) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16Uint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*uint8, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint8PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*uint8, error), list []*int16) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*string, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16StrPtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*string, error), list []*int16) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*bool, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16BoolPtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*bool, error), list []*int16) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16Float32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*float32, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Float32PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*float32, error), list []*int16) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt16Float64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*float64, error)
// 	3. List of type *int16
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Float64PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*float64, error), list []*int16) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*int, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8IntPtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*int, error), list []*int8) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8Int64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*int64, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int64PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*int64, error), list []*int8) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8Int32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*int32, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int32PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*int32, error), list []*int8) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8Int16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*int16, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int16PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*int16, error), list []*int8) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*uint, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8UintPtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*uint, error), list []*int8) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8Uint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*uint64, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint64PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*uint64, error), list []*int8) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8Uint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*uint32, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint32PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*uint32, error), list []*int8) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8Uint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*uint16, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint16PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*uint16, error), list []*int8) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8Uint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*uint8, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint8PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*uint8, error), list []*int8) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*string, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8StrPtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*string, error), list []*int8) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*bool, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8BoolPtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*bool, error), list []*int8) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8Float32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*float32, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Float32PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*float32, error), list []*int8) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapInt8Float64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*float64, error)
// 	3. List of type *int8
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Float64PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*float64, error), list []*int8) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintIntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*int, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintIntPtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*int, error), list []*uint) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintInt64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*int64, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt64PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*int64, error), list []*uint) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintInt32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*int32, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt32PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*int32, error), list []*uint) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintInt16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*int16, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt16PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*int16, error), list []*uint) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintInt8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*int8, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt8PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*int8, error), list []*uint) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintUint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*uint64, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint64PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*uint64, error), list []*uint) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintUint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*uint32, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint32PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*uint32, error), list []*uint) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintUint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*uint16, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint16PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*uint16, error), list []*uint) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintUint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*uint8, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint8PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*uint8, error), list []*uint) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintStrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*string, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintStrPtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*string, error), list []*uint) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintBoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*bool, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintBoolPtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*bool, error), list []*uint) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintFloat32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*float32, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintFloat32PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*float32, error), list []*uint) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUintFloat64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*float64, error)
// 	3. List of type *uint
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintFloat64PtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*float64, error), list []*uint) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*int, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64IntPtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*int, error), list []*uint64) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64Int64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*int64, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int64PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*int64, error), list []*uint64) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64Int32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*int32, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int32PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*int32, error), list []*uint64) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64Int16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*int16, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int16PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*int16, error), list []*uint64) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64Int8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*int8, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int8PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*int8, error), list []*uint64) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*uint, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64UintPtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*uint, error), list []*uint64) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64Uint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*uint32, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint32PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*uint32, error), list []*uint64) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64Uint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*uint16, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint16PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*uint16, error), list []*uint64) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64Uint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*uint8, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint8PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*uint8, error), list []*uint64) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*string, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64StrPtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*string, error), list []*uint64) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*bool, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64BoolPtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*bool, error), list []*uint64) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64Float32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*float32, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Float32PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*float32, error), list []*uint64) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint64Float64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*float64, error)
// 	3. List of type *uint64
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Float64PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*float64, error), list []*uint64) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*int, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32IntPtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*int, error), list []*uint32) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32Int64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*int64, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int64PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*int64, error), list []*uint32) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32Int32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*int32, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int32PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*int32, error), list []*uint32) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32Int16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*int16, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int16PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*int16, error), list []*uint32) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32Int8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*int8, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int8PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*int8, error), list []*uint32) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*uint, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32UintPtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*uint, error), list []*uint32) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32Uint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*uint64, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint64PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*uint64, error), list []*uint32) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32Uint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*uint16, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint16PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*uint16, error), list []*uint32) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32Uint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*uint8, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint8PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*uint8, error), list []*uint32) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*string, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32StrPtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*string, error), list []*uint32) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*bool, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32BoolPtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*bool, error), list []*uint32) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32Float32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*float32, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Float32PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*float32, error), list []*uint32) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint32Float64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*float64, error)
// 	3. List of type *uint32
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Float64PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*float64, error), list []*uint32) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*int, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16IntPtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*int, error), list []*uint16) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16Int64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*int64, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int64PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*int64, error), list []*uint16) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16Int32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*int32, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int32PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*int32, error), list []*uint16) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16Int16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*int16, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int16PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*int16, error), list []*uint16) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16Int8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*int8, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int8PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*int8, error), list []*uint16) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*uint, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16UintPtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*uint, error), list []*uint16) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16Uint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*uint64, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint64PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*uint64, error), list []*uint16) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16Uint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*uint32, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint32PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*uint32, error), list []*uint16) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16Uint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*uint8, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint8PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*uint8, error), list []*uint16) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*string, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16StrPtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*string, error), list []*uint16) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*bool, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16BoolPtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*bool, error), list []*uint16) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16Float32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*float32, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Float32PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*float32, error), list []*uint16) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint16Float64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*float64, error)
// 	3. List of type *uint16
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Float64PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*float64, error), list []*uint16) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*int, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8IntPtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*int, error), list []*uint8) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8Int64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*int64, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int64PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*int64, error), list []*uint8) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8Int32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*int32, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int32PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*int32, error), list []*uint8) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8Int16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*int16, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int16PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*int16, error), list []*uint8) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8Int8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*int8, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int8PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*int8, error), list []*uint8) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*uint, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8UintPtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*uint, error), list []*uint8) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8Uint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*uint64, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint64PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*uint64, error), list []*uint8) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8Uint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*uint32, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint32PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*uint32, error), list []*uint8) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8Uint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*uint16, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint16PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*uint16, error), list []*uint8) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*string, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8StrPtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*string, error), list []*uint8) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*bool, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8BoolPtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*bool, error), list []*uint8) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8Float32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*float32, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Float32PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*float32, error), list []*uint8) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapUint8Float64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*float64, error)
// 	3. List of type *uint8
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Float64PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*float64, error), list []*uint8) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrIntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*int, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrIntPtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*int, error), list []*string) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrInt64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*int64, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt64PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*int64, error), list []*string) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrInt32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*int32, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt32PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*int32, error), list []*string) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrInt16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*int16, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt16PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*int16, error), list []*string) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrInt8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*int8, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt8PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*int8, error), list []*string) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrUintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*uint, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUintPtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*uint, error), list []*string) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrUint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*uint64, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint64PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*uint64, error), list []*string) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrUint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*uint32, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint32PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*uint32, error), list []*string) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrUint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*uint16, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint16PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*uint16, error), list []*string) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrUint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*uint8, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint8PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*uint8, error), list []*string) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrBoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*bool, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrBoolPtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*bool, error), list []*string) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrFloat32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*float32, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrFloat32PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*float32, error), list []*string) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapStrFloat64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns (bool, error).
//	2. Function: takes *string as argument and returns (*float64, error)
// 	3. List of type *string
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrFloat64PtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*float64, error), list []*string) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolIntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*int, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolIntPtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*int, error), list []*bool) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolInt64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*int64, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt64PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*int64, error), list []*bool) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolInt32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*int32, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt32PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*int32, error), list []*bool) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolInt16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*int16, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt16PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*int16, error), list []*bool) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolInt8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*int8, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt8PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*int8, error), list []*bool) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolUintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*uint, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUintPtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*uint, error), list []*bool) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolUint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*uint64, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint64PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*uint64, error), list []*bool) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolUint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*uint32, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint32PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*uint32, error), list []*bool) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolUint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*uint16, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint16PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*uint16, error), list []*bool) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolUint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*uint8, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint8PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*uint8, error), list []*bool) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolStrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*string, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolStrPtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*string, error), list []*bool) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolFloat32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*float32, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolFloat32PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*float32, error), list []*bool) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapBoolFloat64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns (bool, error).
//	2. Function: takes *bool as argument and returns (*float64, error)
// 	3. List of type *bool
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolFloat64PtrErr(fFilter func(*bool) (bool, error), fMap func(*bool) (*float64, error), list []*bool) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*int, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32IntPtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*int, error), list []*float32) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32Int64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*int64, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int64PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*int64, error), list []*float32) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32Int32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*int32, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int32PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*int32, error), list []*float32) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32Int16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*int16, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int16PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*int16, error), list []*float32) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32Int8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*int8, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int8PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*int8, error), list []*float32) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*uint, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32UintPtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*uint, error), list []*float32) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32Uint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*uint64, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint64PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*uint64, error), list []*float32) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32Uint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*uint32, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint32PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*uint32, error), list []*float32) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32Uint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*uint16, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint16PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*uint16, error), list []*float32) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32Uint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*uint8, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint8PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*uint8, error), list []*float32) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*string, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32StrPtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*string, error), list []*float32) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*bool, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32BoolPtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*bool, error), list []*float32) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat32Float64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*float64, error)
// 	3. List of type *float32
//
// Returns:
//	New List of type *float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Float64PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*float64, error), list []*float32) ([]*float64, error) {
	if fFilter == nil || fMap == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64IntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*int, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64IntPtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*int, error), list []*float64) ([]*int, error) {
	if fFilter == nil || fMap == nil {
		return []*int{}, nil
	}
	var newList []*int
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64Int64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*int64, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int64PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*int64, error), list []*float64) ([]*int64, error) {
	if fFilter == nil || fMap == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64Int32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*int32, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int32PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*int32, error), list []*float64) ([]*int32, error) {
	if fFilter == nil || fMap == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64Int16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*int16, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int16PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*int16, error), list []*float64) ([]*int16, error) {
	if fFilter == nil || fMap == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64Int8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*int8, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int8PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*int8, error), list []*float64) ([]*int8, error) {
	if fFilter == nil || fMap == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64UintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*uint, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64UintPtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*uint, error), list []*float64) ([]*uint, error) {
	if fFilter == nil || fMap == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64Uint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*uint64, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint64PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*uint64, error), list []*float64) ([]*uint64, error) {
	if fFilter == nil || fMap == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64Uint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*uint32, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint32PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*uint32, error), list []*float64) ([]*uint32, error) {
	if fFilter == nil || fMap == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64Uint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*uint16, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint16PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*uint16, error), list []*float64) ([]*uint16, error) {
	if fFilter == nil || fMap == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64Uint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*uint8, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint8PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*uint8, error), list []*float64) ([]*uint8, error) {
	if fFilter == nil || fMap == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64StrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*string, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64StrPtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*string, error), list []*float64) ([]*string, error) {
	if fFilter == nil || fMap == nil {
		return []*string{}, nil
	}
	var newList []*string
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64BoolPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*bool, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64BoolPtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*bool, error), list []*float64) ([]*bool, error) {
	if fFilter == nil || fMap == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}

// FilterMapFloat64Float32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*float32, error)
// 	3. List of type *float64
//
// Returns:
//	New List of type *float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Float32PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*float32, error), list []*float64) ([]*float32, error) {
	if fFilter == nil || fMap == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for _, v := range list {
		r, err := fFilter(v)
		if err != nil {
			return nil, err
		}
		if r {
			r, err := fMap(v)
			if err != nil {
				return nil, err
			}
			newList = append(newList, r)
		}
	}
	return newList, nil
}
