package fp

// FilterMapIntInt64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (int64, error)
// 	3. List of type int
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt64Err(fFilter func(int) (bool, error), fMap func(int) (int64, error), list []int) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapIntInt32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (int32, error)
// 	3. List of type int
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt32Err(fFilter func(int) (bool, error), fMap func(int) (int32, error), list []int) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapIntInt16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (int16, error)
// 	3. List of type int
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt16Err(fFilter func(int) (bool, error), fMap func(int) (int16, error), list []int) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapIntInt8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (int8, error)
// 	3. List of type int
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt8Err(fFilter func(int) (bool, error), fMap func(int) (int8, error), list []int) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapIntUintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (uint, error)
// 	3. List of type int
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUintErr(fFilter func(int) (bool, error), fMap func(int) (uint, error), list []int) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapIntUint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (uint64, error)
// 	3. List of type int
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint64Err(fFilter func(int) (bool, error), fMap func(int) (uint64, error), list []int) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapIntUint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (uint32, error)
// 	3. List of type int
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint32Err(fFilter func(int) (bool, error), fMap func(int) (uint32, error), list []int) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapIntUint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (uint16, error)
// 	3. List of type int
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint16Err(fFilter func(int) (bool, error), fMap func(int) (uint16, error), list []int) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapIntUint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (uint8, error)
// 	3. List of type int
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint8Err(fFilter func(int) (bool, error), fMap func(int) (uint8, error), list []int) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapIntStrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (string, error)
// 	3. List of type int
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntStrErr(fFilter func(int) (bool, error), fMap func(int) (string, error), list []int) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapIntBoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (bool, error)
// 	3. List of type int
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntBoolErr(fFilter func(int) (bool, error), fMap func(int) (bool, error), list []int) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapIntFloat32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (float32, error)
// 	3. List of type int
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntFloat32Err(fFilter func(int) (bool, error), fMap func(int) (float32, error), list []int) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapIntFloat64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int and returns (bool, error).
//	2. Function: takes int as argument and returns (float64, error)
// 	3. List of type int
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntFloat64Err(fFilter func(int) (bool, error), fMap func(int) (float64, error), list []int) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapInt64IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (int, error)
// 	3. List of type int64
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64IntErr(fFilter func(int64) (bool, error), fMap func(int64) (int, error), list []int64) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapInt64Int32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (int32, error)
// 	3. List of type int64
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int32Err(fFilter func(int64) (bool, error), fMap func(int64) (int32, error), list []int64) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapInt64Int16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (int16, error)
// 	3. List of type int64
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int16Err(fFilter func(int64) (bool, error), fMap func(int64) (int16, error), list []int64) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapInt64Int8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (int8, error)
// 	3. List of type int64
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int8Err(fFilter func(int64) (bool, error), fMap func(int64) (int8, error), list []int64) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapInt64UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (uint, error)
// 	3. List of type int64
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64UintErr(fFilter func(int64) (bool, error), fMap func(int64) (uint, error), list []int64) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapInt64Uint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (uint64, error)
// 	3. List of type int64
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint64Err(fFilter func(int64) (bool, error), fMap func(int64) (uint64, error), list []int64) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapInt64Uint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (uint32, error)
// 	3. List of type int64
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint32Err(fFilter func(int64) (bool, error), fMap func(int64) (uint32, error), list []int64) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapInt64Uint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (uint16, error)
// 	3. List of type int64
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint16Err(fFilter func(int64) (bool, error), fMap func(int64) (uint16, error), list []int64) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapInt64Uint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (uint8, error)
// 	3. List of type int64
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint8Err(fFilter func(int64) (bool, error), fMap func(int64) (uint8, error), list []int64) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapInt64StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (string, error)
// 	3. List of type int64
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64StrErr(fFilter func(int64) (bool, error), fMap func(int64) (string, error), list []int64) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapInt64BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (bool, error)
// 	3. List of type int64
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64BoolErr(fFilter func(int64) (bool, error), fMap func(int64) (bool, error), list []int64) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapInt64Float32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (float32, error)
// 	3. List of type int64
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Float32Err(fFilter func(int64) (bool, error), fMap func(int64) (float32, error), list []int64) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapInt64Float64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns (bool, error).
//	2. Function: takes int64 as argument and returns (float64, error)
// 	3. List of type int64
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Float64Err(fFilter func(int64) (bool, error), fMap func(int64) (float64, error), list []int64) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapInt32IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (int, error)
// 	3. List of type int32
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32IntErr(fFilter func(int32) (bool, error), fMap func(int32) (int, error), list []int32) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapInt32Int64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (int64, error)
// 	3. List of type int32
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int64Err(fFilter func(int32) (bool, error), fMap func(int32) (int64, error), list []int32) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapInt32Int16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (int16, error)
// 	3. List of type int32
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int16Err(fFilter func(int32) (bool, error), fMap func(int32) (int16, error), list []int32) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapInt32Int8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (int8, error)
// 	3. List of type int32
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int8Err(fFilter func(int32) (bool, error), fMap func(int32) (int8, error), list []int32) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapInt32UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (uint, error)
// 	3. List of type int32
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32UintErr(fFilter func(int32) (bool, error), fMap func(int32) (uint, error), list []int32) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapInt32Uint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (uint64, error)
// 	3. List of type int32
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint64Err(fFilter func(int32) (bool, error), fMap func(int32) (uint64, error), list []int32) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapInt32Uint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (uint32, error)
// 	3. List of type int32
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint32Err(fFilter func(int32) (bool, error), fMap func(int32) (uint32, error), list []int32) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapInt32Uint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (uint16, error)
// 	3. List of type int32
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint16Err(fFilter func(int32) (bool, error), fMap func(int32) (uint16, error), list []int32) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapInt32Uint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (uint8, error)
// 	3. List of type int32
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint8Err(fFilter func(int32) (bool, error), fMap func(int32) (uint8, error), list []int32) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapInt32StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (string, error)
// 	3. List of type int32
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32StrErr(fFilter func(int32) (bool, error), fMap func(int32) (string, error), list []int32) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapInt32BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (bool, error)
// 	3. List of type int32
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32BoolErr(fFilter func(int32) (bool, error), fMap func(int32) (bool, error), list []int32) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapInt32Float32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (float32, error)
// 	3. List of type int32
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Float32Err(fFilter func(int32) (bool, error), fMap func(int32) (float32, error), list []int32) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapInt32Float64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns (bool, error).
//	2. Function: takes int32 as argument and returns (float64, error)
// 	3. List of type int32
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Float64Err(fFilter func(int32) (bool, error), fMap func(int32) (float64, error), list []int32) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapInt16IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (int, error)
// 	3. List of type int16
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16IntErr(fFilter func(int16) (bool, error), fMap func(int16) (int, error), list []int16) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapInt16Int64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (int64, error)
// 	3. List of type int16
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int64Err(fFilter func(int16) (bool, error), fMap func(int16) (int64, error), list []int16) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapInt16Int32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (int32, error)
// 	3. List of type int16
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int32Err(fFilter func(int16) (bool, error), fMap func(int16) (int32, error), list []int16) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapInt16Int8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (int8, error)
// 	3. List of type int16
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int8Err(fFilter func(int16) (bool, error), fMap func(int16) (int8, error), list []int16) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapInt16UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (uint, error)
// 	3. List of type int16
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16UintErr(fFilter func(int16) (bool, error), fMap func(int16) (uint, error), list []int16) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapInt16Uint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (uint64, error)
// 	3. List of type int16
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint64Err(fFilter func(int16) (bool, error), fMap func(int16) (uint64, error), list []int16) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapInt16Uint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (uint32, error)
// 	3. List of type int16
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint32Err(fFilter func(int16) (bool, error), fMap func(int16) (uint32, error), list []int16) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapInt16Uint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (uint16, error)
// 	3. List of type int16
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint16Err(fFilter func(int16) (bool, error), fMap func(int16) (uint16, error), list []int16) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapInt16Uint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (uint8, error)
// 	3. List of type int16
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint8Err(fFilter func(int16) (bool, error), fMap func(int16) (uint8, error), list []int16) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapInt16StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (string, error)
// 	3. List of type int16
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16StrErr(fFilter func(int16) (bool, error), fMap func(int16) (string, error), list []int16) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapInt16BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (bool, error)
// 	3. List of type int16
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16BoolErr(fFilter func(int16) (bool, error), fMap func(int16) (bool, error), list []int16) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapInt16Float32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (float32, error)
// 	3. List of type int16
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Float32Err(fFilter func(int16) (bool, error), fMap func(int16) (float32, error), list []int16) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapInt16Float64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns (bool, error).
//	2. Function: takes int16 as argument and returns (float64, error)
// 	3. List of type int16
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Float64Err(fFilter func(int16) (bool, error), fMap func(int16) (float64, error), list []int16) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapInt8IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (int, error)
// 	3. List of type int8
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8IntErr(fFilter func(int8) (bool, error), fMap func(int8) (int, error), list []int8) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapInt8Int64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (int64, error)
// 	3. List of type int8
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int64Err(fFilter func(int8) (bool, error), fMap func(int8) (int64, error), list []int8) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapInt8Int32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (int32, error)
// 	3. List of type int8
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int32Err(fFilter func(int8) (bool, error), fMap func(int8) (int32, error), list []int8) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapInt8Int16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (int16, error)
// 	3. List of type int8
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int16Err(fFilter func(int8) (bool, error), fMap func(int8) (int16, error), list []int8) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapInt8UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (uint, error)
// 	3. List of type int8
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8UintErr(fFilter func(int8) (bool, error), fMap func(int8) (uint, error), list []int8) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapInt8Uint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (uint64, error)
// 	3. List of type int8
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint64Err(fFilter func(int8) (bool, error), fMap func(int8) (uint64, error), list []int8) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapInt8Uint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (uint32, error)
// 	3. List of type int8
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint32Err(fFilter func(int8) (bool, error), fMap func(int8) (uint32, error), list []int8) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapInt8Uint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (uint16, error)
// 	3. List of type int8
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint16Err(fFilter func(int8) (bool, error), fMap func(int8) (uint16, error), list []int8) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapInt8Uint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (uint8, error)
// 	3. List of type int8
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint8Err(fFilter func(int8) (bool, error), fMap func(int8) (uint8, error), list []int8) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapInt8StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (string, error)
// 	3. List of type int8
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8StrErr(fFilter func(int8) (bool, error), fMap func(int8) (string, error), list []int8) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapInt8BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (bool, error)
// 	3. List of type int8
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8BoolErr(fFilter func(int8) (bool, error), fMap func(int8) (bool, error), list []int8) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapInt8Float32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (float32, error)
// 	3. List of type int8
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Float32Err(fFilter func(int8) (bool, error), fMap func(int8) (float32, error), list []int8) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapInt8Float64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns (bool, error).
//	2. Function: takes int8 as argument and returns (float64, error)
// 	3. List of type int8
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Float64Err(fFilter func(int8) (bool, error), fMap func(int8) (float64, error), list []int8) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapUintIntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (int, error)
// 	3. List of type uint
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintIntErr(fFilter func(uint) (bool, error), fMap func(uint) (int, error), list []uint) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapUintInt64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (int64, error)
// 	3. List of type uint
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt64Err(fFilter func(uint) (bool, error), fMap func(uint) (int64, error), list []uint) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapUintInt32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (int32, error)
// 	3. List of type uint
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt32Err(fFilter func(uint) (bool, error), fMap func(uint) (int32, error), list []uint) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapUintInt16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (int16, error)
// 	3. List of type uint
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt16Err(fFilter func(uint) (bool, error), fMap func(uint) (int16, error), list []uint) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapUintInt8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (int8, error)
// 	3. List of type uint
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt8Err(fFilter func(uint) (bool, error), fMap func(uint) (int8, error), list []uint) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapUintUint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (uint64, error)
// 	3. List of type uint
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint64Err(fFilter func(uint) (bool, error), fMap func(uint) (uint64, error), list []uint) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapUintUint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (uint32, error)
// 	3. List of type uint
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint32Err(fFilter func(uint) (bool, error), fMap func(uint) (uint32, error), list []uint) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapUintUint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (uint16, error)
// 	3. List of type uint
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint16Err(fFilter func(uint) (bool, error), fMap func(uint) (uint16, error), list []uint) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapUintUint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (uint8, error)
// 	3. List of type uint
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint8Err(fFilter func(uint) (bool, error), fMap func(uint) (uint8, error), list []uint) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapUintStrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (string, error)
// 	3. List of type uint
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintStrErr(fFilter func(uint) (bool, error), fMap func(uint) (string, error), list []uint) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapUintBoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (bool, error)
// 	3. List of type uint
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintBoolErr(fFilter func(uint) (bool, error), fMap func(uint) (bool, error), list []uint) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapUintFloat32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (float32, error)
// 	3. List of type uint
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintFloat32Err(fFilter func(uint) (bool, error), fMap func(uint) (float32, error), list []uint) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapUintFloat64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns (bool, error).
//	2. Function: takes uint as argument and returns (float64, error)
// 	3. List of type uint
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintFloat64Err(fFilter func(uint) (bool, error), fMap func(uint) (float64, error), list []uint) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapUint64IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (int, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64IntErr(fFilter func(uint64) (bool, error), fMap func(uint64) (int, error), list []uint64) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapUint64Int64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (int64, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int64Err(fFilter func(uint64) (bool, error), fMap func(uint64) (int64, error), list []uint64) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapUint64Int32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (int32, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int32Err(fFilter func(uint64) (bool, error), fMap func(uint64) (int32, error), list []uint64) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapUint64Int16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (int16, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int16Err(fFilter func(uint64) (bool, error), fMap func(uint64) (int16, error), list []uint64) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapUint64Int8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (int8, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int8Err(fFilter func(uint64) (bool, error), fMap func(uint64) (int8, error), list []uint64) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapUint64UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (uint, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64UintErr(fFilter func(uint64) (bool, error), fMap func(uint64) (uint, error), list []uint64) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapUint64Uint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (uint32, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint32Err(fFilter func(uint64) (bool, error), fMap func(uint64) (uint32, error), list []uint64) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapUint64Uint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (uint16, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint16Err(fFilter func(uint64) (bool, error), fMap func(uint64) (uint16, error), list []uint64) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapUint64Uint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (uint8, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint8Err(fFilter func(uint64) (bool, error), fMap func(uint64) (uint8, error), list []uint64) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapUint64StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (string, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64StrErr(fFilter func(uint64) (bool, error), fMap func(uint64) (string, error), list []uint64) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapUint64BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (bool, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64BoolErr(fFilter func(uint64) (bool, error), fMap func(uint64) (bool, error), list []uint64) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapUint64Float32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (float32, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Float32Err(fFilter func(uint64) (bool, error), fMap func(uint64) (float32, error), list []uint64) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapUint64Float64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns (bool, error).
//	2. Function: takes uint64 as argument and returns (float64, error)
// 	3. List of type uint64
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Float64Err(fFilter func(uint64) (bool, error), fMap func(uint64) (float64, error), list []uint64) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapUint32IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (int, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32IntErr(fFilter func(uint32) (bool, error), fMap func(uint32) (int, error), list []uint32) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapUint32Int64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (int64, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int64Err(fFilter func(uint32) (bool, error), fMap func(uint32) (int64, error), list []uint32) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapUint32Int32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (int32, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int32Err(fFilter func(uint32) (bool, error), fMap func(uint32) (int32, error), list []uint32) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapUint32Int16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (int16, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int16Err(fFilter func(uint32) (bool, error), fMap func(uint32) (int16, error), list []uint32) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapUint32Int8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (int8, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int8Err(fFilter func(uint32) (bool, error), fMap func(uint32) (int8, error), list []uint32) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapUint32UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (uint, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32UintErr(fFilter func(uint32) (bool, error), fMap func(uint32) (uint, error), list []uint32) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapUint32Uint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (uint64, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint64Err(fFilter func(uint32) (bool, error), fMap func(uint32) (uint64, error), list []uint32) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapUint32Uint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (uint16, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint16Err(fFilter func(uint32) (bool, error), fMap func(uint32) (uint16, error), list []uint32) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapUint32Uint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (uint8, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint8Err(fFilter func(uint32) (bool, error), fMap func(uint32) (uint8, error), list []uint32) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapUint32StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (string, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32StrErr(fFilter func(uint32) (bool, error), fMap func(uint32) (string, error), list []uint32) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapUint32BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (bool, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32BoolErr(fFilter func(uint32) (bool, error), fMap func(uint32) (bool, error), list []uint32) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapUint32Float32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (float32, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Float32Err(fFilter func(uint32) (bool, error), fMap func(uint32) (float32, error), list []uint32) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapUint32Float64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns (bool, error).
//	2. Function: takes uint32 as argument and returns (float64, error)
// 	3. List of type uint32
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Float64Err(fFilter func(uint32) (bool, error), fMap func(uint32) (float64, error), list []uint32) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapUint16IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (int, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16IntErr(fFilter func(uint16) (bool, error), fMap func(uint16) (int, error), list []uint16) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapUint16Int64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (int64, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int64Err(fFilter func(uint16) (bool, error), fMap func(uint16) (int64, error), list []uint16) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapUint16Int32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (int32, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int32Err(fFilter func(uint16) (bool, error), fMap func(uint16) (int32, error), list []uint16) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapUint16Int16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (int16, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int16Err(fFilter func(uint16) (bool, error), fMap func(uint16) (int16, error), list []uint16) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapUint16Int8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (int8, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int8Err(fFilter func(uint16) (bool, error), fMap func(uint16) (int8, error), list []uint16) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapUint16UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (uint, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16UintErr(fFilter func(uint16) (bool, error), fMap func(uint16) (uint, error), list []uint16) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapUint16Uint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (uint64, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint64Err(fFilter func(uint16) (bool, error), fMap func(uint16) (uint64, error), list []uint16) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapUint16Uint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (uint32, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint32Err(fFilter func(uint16) (bool, error), fMap func(uint16) (uint32, error), list []uint16) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapUint16Uint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (uint8, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint8Err(fFilter func(uint16) (bool, error), fMap func(uint16) (uint8, error), list []uint16) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapUint16StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (string, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16StrErr(fFilter func(uint16) (bool, error), fMap func(uint16) (string, error), list []uint16) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapUint16BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (bool, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16BoolErr(fFilter func(uint16) (bool, error), fMap func(uint16) (bool, error), list []uint16) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapUint16Float32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (float32, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Float32Err(fFilter func(uint16) (bool, error), fMap func(uint16) (float32, error), list []uint16) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapUint16Float64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns (bool, error).
//	2. Function: takes uint16 as argument and returns (float64, error)
// 	3. List of type uint16
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Float64Err(fFilter func(uint16) (bool, error), fMap func(uint16) (float64, error), list []uint16) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapUint8IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (int, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8IntErr(fFilter func(uint8) (bool, error), fMap func(uint8) (int, error), list []uint8) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapUint8Int64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (int64, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int64Err(fFilter func(uint8) (bool, error), fMap func(uint8) (int64, error), list []uint8) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapUint8Int32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (int32, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int32Err(fFilter func(uint8) (bool, error), fMap func(uint8) (int32, error), list []uint8) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapUint8Int16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (int16, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int16Err(fFilter func(uint8) (bool, error), fMap func(uint8) (int16, error), list []uint8) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapUint8Int8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (int8, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int8Err(fFilter func(uint8) (bool, error), fMap func(uint8) (int8, error), list []uint8) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapUint8UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (uint, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8UintErr(fFilter func(uint8) (bool, error), fMap func(uint8) (uint, error), list []uint8) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapUint8Uint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (uint64, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint64Err(fFilter func(uint8) (bool, error), fMap func(uint8) (uint64, error), list []uint8) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapUint8Uint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (uint32, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint32Err(fFilter func(uint8) (bool, error), fMap func(uint8) (uint32, error), list []uint8) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapUint8Uint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (uint16, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint16Err(fFilter func(uint8) (bool, error), fMap func(uint8) (uint16, error), list []uint8) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapUint8StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (string, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8StrErr(fFilter func(uint8) (bool, error), fMap func(uint8) (string, error), list []uint8) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapUint8BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (bool, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8BoolErr(fFilter func(uint8) (bool, error), fMap func(uint8) (bool, error), list []uint8) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapUint8Float32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (float32, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Float32Err(fFilter func(uint8) (bool, error), fMap func(uint8) (float32, error), list []uint8) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapUint8Float64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns (bool, error).
//	2. Function: takes uint8 as argument and returns (float64, error)
// 	3. List of type uint8
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Float64Err(fFilter func(uint8) (bool, error), fMap func(uint8) (float64, error), list []uint8) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapStrIntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (int, error)
// 	3. List of type string
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrIntErr(fFilter func(string) (bool, error), fMap func(string) (int, error), list []string) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapStrInt64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (int64, error)
// 	3. List of type string
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt64Err(fFilter func(string) (bool, error), fMap func(string) (int64, error), list []string) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapStrInt32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (int32, error)
// 	3. List of type string
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt32Err(fFilter func(string) (bool, error), fMap func(string) (int32, error), list []string) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapStrInt16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (int16, error)
// 	3. List of type string
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt16Err(fFilter func(string) (bool, error), fMap func(string) (int16, error), list []string) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapStrInt8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (int8, error)
// 	3. List of type string
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt8Err(fFilter func(string) (bool, error), fMap func(string) (int8, error), list []string) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapStrUintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (uint, error)
// 	3. List of type string
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUintErr(fFilter func(string) (bool, error), fMap func(string) (uint, error), list []string) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapStrUint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (uint64, error)
// 	3. List of type string
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint64Err(fFilter func(string) (bool, error), fMap func(string) (uint64, error), list []string) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapStrUint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (uint32, error)
// 	3. List of type string
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint32Err(fFilter func(string) (bool, error), fMap func(string) (uint32, error), list []string) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapStrUint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (uint16, error)
// 	3. List of type string
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint16Err(fFilter func(string) (bool, error), fMap func(string) (uint16, error), list []string) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapStrUint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (uint8, error)
// 	3. List of type string
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint8Err(fFilter func(string) (bool, error), fMap func(string) (uint8, error), list []string) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapStrBoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (bool, error)
// 	3. List of type string
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrBoolErr(fFilter func(string) (bool, error), fMap func(string) (bool, error), list []string) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapStrFloat32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (float32, error)
// 	3. List of type string
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrFloat32Err(fFilter func(string) (bool, error), fMap func(string) (float32, error), list []string) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapStrFloat64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - string and returns (bool, error).
//	2. Function: takes string as argument and returns (float64, error)
// 	3. List of type string
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrFloat64Err(fFilter func(string) (bool, error), fMap func(string) (float64, error), list []string) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapBoolIntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (int, error)
// 	3. List of type bool
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolIntErr(fFilter func(bool) (bool, error), fMap func(bool) (int, error), list []bool) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapBoolInt64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (int64, error)
// 	3. List of type bool
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt64Err(fFilter func(bool) (bool, error), fMap func(bool) (int64, error), list []bool) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapBoolInt32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (int32, error)
// 	3. List of type bool
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt32Err(fFilter func(bool) (bool, error), fMap func(bool) (int32, error), list []bool) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapBoolInt16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (int16, error)
// 	3. List of type bool
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt16Err(fFilter func(bool) (bool, error), fMap func(bool) (int16, error), list []bool) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapBoolInt8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (int8, error)
// 	3. List of type bool
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt8Err(fFilter func(bool) (bool, error), fMap func(bool) (int8, error), list []bool) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapBoolUintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (uint, error)
// 	3. List of type bool
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUintErr(fFilter func(bool) (bool, error), fMap func(bool) (uint, error), list []bool) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapBoolUint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (uint64, error)
// 	3. List of type bool
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint64Err(fFilter func(bool) (bool, error), fMap func(bool) (uint64, error), list []bool) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapBoolUint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (uint32, error)
// 	3. List of type bool
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint32Err(fFilter func(bool) (bool, error), fMap func(bool) (uint32, error), list []bool) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapBoolUint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (uint16, error)
// 	3. List of type bool
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint16Err(fFilter func(bool) (bool, error), fMap func(bool) (uint16, error), list []bool) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapBoolUint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (uint8, error)
// 	3. List of type bool
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint8Err(fFilter func(bool) (bool, error), fMap func(bool) (uint8, error), list []bool) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapBoolStrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (string, error)
// 	3. List of type bool
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolStrErr(fFilter func(bool) (bool, error), fMap func(bool) (string, error), list []bool) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapBoolFloat32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (float32, error)
// 	3. List of type bool
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolFloat32Err(fFilter func(bool) (bool, error), fMap func(bool) (float32, error), list []bool) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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

// FilterMapBoolFloat64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns (bool, error).
//	2. Function: takes bool as argument and returns (float64, error)
// 	3. List of type bool
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolFloat64Err(fFilter func(bool) (bool, error), fMap func(bool) (float64, error), list []bool) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapFloat32IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (int, error)
// 	3. List of type float32
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32IntErr(fFilter func(float32) (bool, error), fMap func(float32) (int, error), list []float32) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapFloat32Int64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (int64, error)
// 	3. List of type float32
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int64Err(fFilter func(float32) (bool, error), fMap func(float32) (int64, error), list []float32) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapFloat32Int32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (int32, error)
// 	3. List of type float32
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int32Err(fFilter func(float32) (bool, error), fMap func(float32) (int32, error), list []float32) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapFloat32Int16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (int16, error)
// 	3. List of type float32
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int16Err(fFilter func(float32) (bool, error), fMap func(float32) (int16, error), list []float32) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapFloat32Int8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (int8, error)
// 	3. List of type float32
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int8Err(fFilter func(float32) (bool, error), fMap func(float32) (int8, error), list []float32) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapFloat32UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (uint, error)
// 	3. List of type float32
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32UintErr(fFilter func(float32) (bool, error), fMap func(float32) (uint, error), list []float32) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapFloat32Uint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (uint64, error)
// 	3. List of type float32
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint64Err(fFilter func(float32) (bool, error), fMap func(float32) (uint64, error), list []float32) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapFloat32Uint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (uint32, error)
// 	3. List of type float32
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint32Err(fFilter func(float32) (bool, error), fMap func(float32) (uint32, error), list []float32) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapFloat32Uint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (uint16, error)
// 	3. List of type float32
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint16Err(fFilter func(float32) (bool, error), fMap func(float32) (uint16, error), list []float32) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapFloat32Uint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (uint8, error)
// 	3. List of type float32
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint8Err(fFilter func(float32) (bool, error), fMap func(float32) (uint8, error), list []float32) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapFloat32StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (string, error)
// 	3. List of type float32
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32StrErr(fFilter func(float32) (bool, error), fMap func(float32) (string, error), list []float32) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapFloat32BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (bool, error)
// 	3. List of type float32
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32BoolErr(fFilter func(float32) (bool, error), fMap func(float32) (bool, error), list []float32) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapFloat32Float64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float32 and returns (bool, error).
//	2. Function: takes float32 as argument and returns (float64, error)
// 	3. List of type float32
//
// Returns:
//	New List of type float64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Float64Err(fFilter func(float32) (bool, error), fMap func(float32) (float64, error), list []float32) ([]float64, error) {
	if fFilter == nil || fMap == nil {
		return []float64{}, nil
	}
	var newList []float64
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

// FilterMapFloat64IntErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (int, error)
// 	3. List of type float64
//
// Returns:
//	New List of type int, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64IntErr(fFilter func(float64) (bool, error), fMap func(float64) (int, error), list []float64) ([]int, error) {
	if fFilter == nil || fMap == nil {
		return []int{}, nil
	}
	var newList []int
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

// FilterMapFloat64Int64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (int64, error)
// 	3. List of type float64
//
// Returns:
//	New List of type int64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int64Err(fFilter func(float64) (bool, error), fMap func(float64) (int64, error), list []float64) ([]int64, error) {
	if fFilter == nil || fMap == nil {
		return []int64{}, nil
	}
	var newList []int64
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

// FilterMapFloat64Int32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (int32, error)
// 	3. List of type float64
//
// Returns:
//	New List of type int32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int32Err(fFilter func(float64) (bool, error), fMap func(float64) (int32, error), list []float64) ([]int32, error) {
	if fFilter == nil || fMap == nil {
		return []int32{}, nil
	}
	var newList []int32
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

// FilterMapFloat64Int16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (int16, error)
// 	3. List of type float64
//
// Returns:
//	New List of type int16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int16Err(fFilter func(float64) (bool, error), fMap func(float64) (int16, error), list []float64) ([]int16, error) {
	if fFilter == nil || fMap == nil {
		return []int16{}, nil
	}
	var newList []int16
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

// FilterMapFloat64Int8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (int8, error)
// 	3. List of type float64
//
// Returns:
//	New List of type int8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int8Err(fFilter func(float64) (bool, error), fMap func(float64) (int8, error), list []float64) ([]int8, error) {
	if fFilter == nil || fMap == nil {
		return []int8{}, nil
	}
	var newList []int8
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

// FilterMapFloat64UintErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (uint, error)
// 	3. List of type float64
//
// Returns:
//	New List of type uint, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64UintErr(fFilter func(float64) (bool, error), fMap func(float64) (uint, error), list []float64) ([]uint, error) {
	if fFilter == nil || fMap == nil {
		return []uint{}, nil
	}
	var newList []uint
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

// FilterMapFloat64Uint64Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (uint64, error)
// 	3. List of type float64
//
// Returns:
//	New List of type uint64, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint64Err(fFilter func(float64) (bool, error), fMap func(float64) (uint64, error), list []float64) ([]uint64, error) {
	if fFilter == nil || fMap == nil {
		return []uint64{}, nil
	}
	var newList []uint64
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

// FilterMapFloat64Uint32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (uint32, error)
// 	3. List of type float64
//
// Returns:
//	New List of type uint32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint32Err(fFilter func(float64) (bool, error), fMap func(float64) (uint32, error), list []float64) ([]uint32, error) {
	if fFilter == nil || fMap == nil {
		return []uint32{}, nil
	}
	var newList []uint32
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

// FilterMapFloat64Uint16Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (uint16, error)
// 	3. List of type float64
//
// Returns:
//	New List of type uint16, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint16Err(fFilter func(float64) (bool, error), fMap func(float64) (uint16, error), list []float64) ([]uint16, error) {
	if fFilter == nil || fMap == nil {
		return []uint16{}, nil
	}
	var newList []uint16
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

// FilterMapFloat64Uint8Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (uint8, error)
// 	3. List of type float64
//
// Returns:
//	New List of type uint8, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint8Err(fFilter func(float64) (bool, error), fMap func(float64) (uint8, error), list []float64) ([]uint8, error) {
	if fFilter == nil || fMap == nil {
		return []uint8{}, nil
	}
	var newList []uint8
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

// FilterMapFloat64StrErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (string, error)
// 	3. List of type float64
//
// Returns:
//	New List of type string, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64StrErr(fFilter func(float64) (bool, error), fMap func(float64) (string, error), list []float64) ([]string, error) {
	if fFilter == nil || fMap == nil {
		return []string{}, nil
	}
	var newList []string
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

// FilterMapFloat64BoolErr filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (bool, error)
// 	3. List of type float64
//
// Returns:
//	New List of type bool, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64BoolErr(fFilter func(float64) (bool, error), fMap func(float64) (bool, error), list []float64) ([]bool, error) {
	if fFilter == nil || fMap == nil {
		return []bool{}, nil
	}
	var newList []bool
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

// FilterMapFloat64Float32Err filters given list, then apply function(2nd argument) on each item in the list and returns a new list, error
// Takes 3 inputs
//	1. Function: takes one input type - float64 and returns (bool, error).
//	2. Function: takes float64 as argument and returns (float32, error)
// 	3. List of type float64
//
// Returns:
//	New List of type float32, error
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Float32Err(fFilter func(float64) (bool, error), fMap func(float64) (float32, error), list []float64) ([]float32, error) {
	if fFilter == nil || fMap == nil {
		return []float32{}, nil
	}
	var newList []float32
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
