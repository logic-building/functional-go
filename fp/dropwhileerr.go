package fp

// DropWhileIntErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileIntErr(f func(int) (bool, error), list []int) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	var newList []int
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]int, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileInt64Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt64Err(f func(int64) (bool, error), list []int64) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	var newList []int64
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]int64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileInt32Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt32Err(f func(int32) (bool, error), list []int32) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	var newList []int32
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]int32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileInt16Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt16Err(f func(int16) (bool, error), list []int16) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	var newList []int16
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]int16, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileInt8Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt8Err(f func(int8) (bool, error), list []int8) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	var newList []int8
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]int8, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileUintErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUintErr(f func(uint) (bool, error), list []uint) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	var newList []uint
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]uint, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileUint64Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint64Err(f func(uint64) (bool, error), list []uint64) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	var newList []uint64
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]uint64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileUint32Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint32Err(f func(uint32) (bool, error), list []uint32) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	var newList []uint32
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]uint32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileUint16Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint16Err(f func(uint16) (bool, error), list []uint16) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	var newList []uint16
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]uint16, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileUint8Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint8Err(f func(uint8) (bool, error), list []uint8) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	var newList []uint8
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]uint8, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileStrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileStrErr(f func(string) (bool, error), list []string) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	var newList []string
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]string, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileBoolErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileBoolErr(f func(bool) (bool, error), list []bool) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	var newList []bool
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]bool, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileFloat32Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileFloat32Err(f func(float32) (bool, error), list []float32) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	var newList []float32
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]float32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}

// DropWhileFloat64Err drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileFloat64Err(f func(float64) (bool, error), list []float64) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	var newList []float64
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]float64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList, nil
		}
	}
	return newList, nil
}
