package fp

// TakeWhileIntPtr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileIntPtr(f func(*int) bool, list []*int) []*int {
	if f == nil {
		return []*int{}
	}
	var newList []*int
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileInt64Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileInt64Ptr(f func(*int64) bool, list []*int64) []*int64 {
	if f == nil {
		return []*int64{}
	}
	var newList []*int64
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileInt32Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileInt32Ptr(f func(*int32) bool, list []*int32) []*int32 {
	if f == nil {
		return []*int32{}
	}
	var newList []*int32
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileInt16Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileInt16Ptr(f func(*int16) bool, list []*int16) []*int16 {
	if f == nil {
		return []*int16{}
	}
	var newList []*int16
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileInt8Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileInt8Ptr(f func(*int8) bool, list []*int8) []*int8 {
	if f == nil {
		return []*int8{}
	}
	var newList []*int8
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUintPtr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUintPtr(f func(*uint) bool, list []*uint) []*uint {
	if f == nil {
		return []*uint{}
	}
	var newList []*uint
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUint64Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUint64Ptr(f func(*uint64) bool, list []*uint64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	var newList []*uint64
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUint32Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUint32Ptr(f func(*uint32) bool, list []*uint32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	var newList []*uint32
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUint16Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUint16Ptr(f func(*uint16) bool, list []*uint16) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	var newList []*uint16
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUint8Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUint8Ptr(f func(*uint8) bool, list []*uint8) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	var newList []*uint8
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileStrPtr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileStrPtr(f func(*string) bool, list []*string) []*string {
	if f == nil {
		return []*string{}
	}
	var newList []*string
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileBoolPtr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileBoolPtr(f func(*bool) bool, list []*bool) []*bool {
	if f == nil {
		return []*bool{}
	}
	var newList []*bool
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileFloat32Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileFloat32Ptr(f func(*float32) bool, list []*float32) []*float32 {
	if f == nil {
		return []*float32{}
	}
	var newList []*float32
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileFloat64Ptr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileFloat64Ptr(f func(*float64) bool, list []*float64) []*float64 {
	if f == nil {
		return []*float64{}
	}
	var newList []*float64
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}
