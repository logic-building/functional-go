package fp

// FilterMapIntPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*int) and returns (bool, error).
//	2. Function: takes *int as argument and returns (*int, error)
// 	3. Slice of type []*int
//
// Returns:
//	New List ([]*int, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntPtrErr(fFilter func(*int) (bool, error), fMap func(*int) (*int, error), list []*int) ([]*int, error) {
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

// FilterMapInt64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*int64) and returns (bool, error).
//	2. Function: takes *int64 as argument and returns (*int64, error)
// 	3. Slice of type []*int64
//
// Returns:
//	New List ([]*int64, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64PtrErr(fFilter func(*int64) (bool, error), fMap func(*int64) (*int64, error), list []*int64) ([]*int64, error) {
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

// FilterMapInt32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*int32) and returns (bool, error).
//	2. Function: takes *int32 as argument and returns (*int32, error)
// 	3. Slice of type []*int32
//
// Returns:
//	New List ([]*int32, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32PtrErr(fFilter func(*int32) (bool, error), fMap func(*int32) (*int32, error), list []*int32) ([]*int32, error) {
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

// FilterMapInt16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*int16) and returns (bool, error).
//	2. Function: takes *int16 as argument and returns (*int16, error)
// 	3. Slice of type []*int16
//
// Returns:
//	New List ([]*int16, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16PtrErr(fFilter func(*int16) (bool, error), fMap func(*int16) (*int16, error), list []*int16) ([]*int16, error) {
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

// FilterMapInt8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*int8) and returns (bool, error).
//	2. Function: takes *int8 as argument and returns (*int8, error)
// 	3. Slice of type []*int8
//
// Returns:
//	New List ([]*int8, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8PtrErr(fFilter func(*int8) (bool, error), fMap func(*int8) (*int8, error), list []*int8) ([]*int8, error) {
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

// FilterMapUintPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*uint) and returns (bool, error).
//	2. Function: takes *uint as argument and returns (*uint, error)
// 	3. Slice of type []*uint
//
// Returns:
//	New List ([]*uint, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintPtrErr(fFilter func(*uint) (bool, error), fMap func(*uint) (*uint, error), list []*uint) ([]*uint, error) {
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

// FilterMapUint64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*uint64) and returns (bool, error).
//	2. Function: takes *uint64 as argument and returns (*uint64, error)
// 	3. Slice of type []*uint64
//
// Returns:
//	New List ([]*uint64, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64PtrErr(fFilter func(*uint64) (bool, error), fMap func(*uint64) (*uint64, error), list []*uint64) ([]*uint64, error) {
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

// FilterMapUint32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*uint32) and returns (bool, error).
//	2. Function: takes *uint32 as argument and returns (*uint32, error)
// 	3. Slice of type []*uint32
//
// Returns:
//	New List ([]*uint32, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32PtrErr(fFilter func(*uint32) (bool, error), fMap func(*uint32) (*uint32, error), list []*uint32) ([]*uint32, error) {
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

// FilterMapUint16PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*uint16) and returns (bool, error).
//	2. Function: takes *uint16 as argument and returns (*uint16, error)
// 	3. Slice of type []*uint16
//
// Returns:
//	New List ([]*uint16, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16PtrErr(fFilter func(*uint16) (bool, error), fMap func(*uint16) (*uint16, error), list []*uint16) ([]*uint16, error) {
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

// FilterMapUint8PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*uint8) and returns (bool, error).
//	2. Function: takes *uint8 as argument and returns (*uint8, error)
// 	3. Slice of type []*uint8
//
// Returns:
//	New List ([]*uint8, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8PtrErr(fFilter func(*uint8) (bool, error), fMap func(*uint8) (*uint8, error), list []*uint8) ([]*uint8, error) {
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

// FilterMapStrPtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*string) and returns (bool, error).
//	2. Function: takes *string as argument and returns (*string, error)
// 	3. Slice of type []*string
//
// Returns:
//	New List ([]*string, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrPtrErr(fFilter func(*string) (bool, error), fMap func(*string) (*string, error), list []*string) ([]*string, error) {
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

// FilterMapFloat32PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*float32) and returns (bool, error).
//	2. Function: takes *float32 as argument and returns (*float32, error)
// 	3. Slice of type []*float32
//
// Returns:
//	New List ([]*float32, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32PtrErr(fFilter func(*float32) (bool, error), fMap func(*float32) (*float32, error), list []*float32) ([]*float32, error) {
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

// FilterMapFloat64PtrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input(*float64) and returns (bool, error).
//	2. Function: takes *float64 as argument and returns (*float64, error)
// 	3. Slice of type []*float64
//
// Returns:
//	New List ([]*float64, error).
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64PtrErr(fFilter func(*float64) (bool, error), fMap func(*float64) (*float64, error), list []*float64) ([]*float64, error) {
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
