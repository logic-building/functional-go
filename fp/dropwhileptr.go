package fp

// DropWhileIntPtr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileIntPtr(f func(*int) bool, list []*int) []*int {
	if f == nil {
		return []*int{}
	}
	var newList []*int
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*int, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileInt64Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt64Ptr(f func(*int64) bool, list []*int64) []*int64 {
	if f == nil {
		return []*int64{}
	}
	var newList []*int64
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*int64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileInt32Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt32Ptr(f func(*int32) bool, list []*int32) []*int32 {
	if f == nil {
		return []*int32{}
	}
	var newList []*int32
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*int32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileInt16Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt16Ptr(f func(*int16) bool, list []*int16) []*int16 {
	if f == nil {
		return []*int16{}
	}
	var newList []*int16
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*int16, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileInt8Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileInt8Ptr(f func(*int8) bool, list []*int8) []*int8 {
	if f == nil {
		return []*int8{}
	}
	var newList []*int8
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*int8, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUintPtr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileUintPtr(f func(*uint) bool, list []*uint) []*uint {
	if f == nil {
		return []*uint{}
	}
	var newList []*uint
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*uint, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUint64Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint64Ptr(f func(*uint64) bool, list []*uint64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	var newList []*uint64
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*uint64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUint32Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint32Ptr(f func(*uint32) bool, list []*uint32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	var newList []*uint32
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*uint32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUint16Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint16Ptr(f func(*uint16) bool, list []*uint16) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	var newList []*uint16
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*uint16, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUint8Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileUint8Ptr(f func(*uint8) bool, list []*uint8) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	var newList []*uint8
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*uint8, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileStrPtr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileStrPtr(f func(*string) bool, list []*string) []*string {
	if f == nil {
		return []*string{}
	}
	var newList []*string
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*string, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileBoolPtr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileBoolPtr(f func(*bool) bool, list []*bool) []*bool {
	if f == nil {
		return []*bool{}
	}
	var newList []*bool
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*bool, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileFloat32Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileFloat32Ptr(f func(*float32) bool, list []*float32) []*float32 {
	if f == nil {
		return []*float32{}
	}
	var newList []*float32
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*float32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileFloat64Ptr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
func DropWhileFloat64Ptr(f func(*float64) bool, list []*float64) []*float64 {
	if f == nil {
		return []*float64{}
	}
	var newList []*float64
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]*float64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}
