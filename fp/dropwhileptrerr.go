package fp

// DropWhileIntPtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileIntPtrErr(f func(*int) (bool, error), list []*int) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	var newList []*int
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*int, listLen-i)
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

// DropWhileInt64PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt64PtrErr(f func(*int64) (bool, error), list []*int64) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	var newList []*int64
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*int64, listLen-i)
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

// DropWhileInt32PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt32PtrErr(f func(*int32) (bool, error), list []*int32) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	var newList []*int32
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*int32, listLen-i)
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

// DropWhileInt16PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt16PtrErr(f func(*int16) (bool, error), list []*int16) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	var newList []*int16
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*int16, listLen-i)
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

// DropWhileInt8PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt8PtrErr(f func(*int8) (bool, error), list []*int8) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	var newList []*int8
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*int8, listLen-i)
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

// DropWhileUintPtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUintPtrErr(f func(*uint) (bool, error), list []*uint) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	var newList []*uint
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*uint, listLen-i)
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

// DropWhileUint64PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint64PtrErr(f func(*uint64) (bool, error), list []*uint64) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	var newList []*uint64
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*uint64, listLen-i)
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

// DropWhileUint32PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint32PtrErr(f func(*uint32) (bool, error), list []*uint32) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	var newList []*uint32
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*uint32, listLen-i)
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

// DropWhileUint16PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint16PtrErr(f func(*uint16) (bool, error), list []*uint16) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	var newList []*uint16
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*uint16, listLen-i)
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

// DropWhileUint8PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint8PtrErr(f func(*uint8) (bool, error), list []*uint8) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	var newList []*uint8
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*uint8, listLen-i)
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

// DropWhileStrPtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileStrPtrErr(f func(*string) (bool, error), list []*string) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	var newList []*string
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*string, listLen-i)
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

// DropWhileBoolPtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileBoolPtrErr(f func(*bool) (bool, error), list []*bool) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	var newList []*bool
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*bool, listLen-i)
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

// DropWhileFloat32PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileFloat32PtrErr(f func(*float32) (bool, error), list []*float32) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	var newList []*float32
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*float32, listLen-i)
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

// DropWhileFloat64PtrErr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns (boolean, error)
//	2. list
//
// Returns:
// 	New List, error
//  Empty list if either one of arguments or both of them are nil
func DropWhileFloat64PtrErr(f func(*float64) (bool, error), list []*float64) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	var newList []*float64
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			listLen := len(list)
			newList = make([]*float64, listLen-i)
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
