package fp

// RemoveIntPtr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveIntPtr(f func(*int) bool, list []*int) []*int {
	if f == nil {
		return []*int{}
	}
	var newList []*int
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveInt64Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveInt64Ptr(f func(*int64) bool, list []*int64) []*int64 {
	if f == nil {
		return []*int64{}
	}
	var newList []*int64
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveInt32Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveInt32Ptr(f func(*int32) bool, list []*int32) []*int32 {
	if f == nil {
		return []*int32{}
	}
	var newList []*int32
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveInt16Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveInt16Ptr(f func(*int16) bool, list []*int16) []*int16 {
	if f == nil {
		return []*int16{}
	}
	var newList []*int16
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveInt8Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveInt8Ptr(f func(*int8) bool, list []*int8) []*int8 {
	if f == nil {
		return []*int8{}
	}
	var newList []*int8
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUintPtr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUintPtr(f func(*uint) bool, list []*uint) []*uint {
	if f == nil {
		return []*uint{}
	}
	var newList []*uint
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUint64Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUint64Ptr(f func(*uint64) bool, list []*uint64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	var newList []*uint64
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUint32Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUint32Ptr(f func(*uint32) bool, list []*uint32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	var newList []*uint32
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUint16Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUint16Ptr(f func(*uint16) bool, list []*uint16) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	var newList []*uint16
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUint8Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveUint8Ptr(f func(*uint8) bool, list []*uint8) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	var newList []*uint8
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveStrPtr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveStrPtr(f func(*string) bool, list []*string) []*string {
	if f == nil {
		return []*string{}
	}
	var newList []*string
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveBoolPtr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveBoolPtr(f func(*bool) bool, list []*bool) []*bool {
	if f == nil {
		return []*bool{}
	}
	var newList []*bool
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveFloat32Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveFloat32Ptr(f func(*float32) bool, list []*float32) []*float32 {
	if f == nil {
		return []*float32{}
	}
	var newList []*float32
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveFloat64Ptr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveFloat64Ptr(f func(*float64) bool, list []*float64) []*float64 {
	if f == nil {
		return []*float64{}
	}
	var newList []*float64
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
