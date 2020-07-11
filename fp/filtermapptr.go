package fp

// FilterMapIntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes int as argument and returns (int)
// 	3. Slice of type []*int
//
// Returns:
//	New List of type []*int.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntPtr(fFilter func(*int) bool, fMap func(*int) *int, list []*int) []*int {
	if fFilter == nil || fMap == nil {
		return []*int{}
	}
	var newList []*int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapInt64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes int64 as argument and returns (int64)
// 	3. Slice of type []*int64
//
// Returns:
//	New List of type []*int64.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Ptr(fFilter func(*int64) bool, fMap func(*int64) *int64, list []*int64) []*int64 {
	if fFilter == nil || fMap == nil {
		return []*int64{}
	}
	var newList []*int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapInt32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes int32 as argument and returns (int32)
// 	3. Slice of type []*int32
//
// Returns:
//	New List of type []*int32.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Ptr(fFilter func(*int32) bool, fMap func(*int32) *int32, list []*int32) []*int32 {
	if fFilter == nil || fMap == nil {
		return []*int32{}
	}
	var newList []*int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapInt16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes int16 as argument and returns (int16)
// 	3. Slice of type []*int16
//
// Returns:
//	New List of type []*int16.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Ptr(fFilter func(*int16) bool, fMap func(*int16) *int16, list []*int16) []*int16 {
	if fFilter == nil || fMap == nil {
		return []*int16{}
	}
	var newList []*int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapInt8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes int8 as argument and returns (int8)
// 	3. Slice of type []*int8
//
// Returns:
//	New List of type []*int8.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Ptr(fFilter func(*int8) bool, fMap func(*int8) *int8, list []*int8) []*int8 {
	if fFilter == nil || fMap == nil {
		return []*int8{}
	}
	var newList []*int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes uint as argument and returns (uint)
// 	3. Slice of type []*uint
//
// Returns:
//	New List of type []*uint.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintPtr(fFilter func(*uint) bool, fMap func(*uint) *uint, list []*uint) []*uint {
	if fFilter == nil || fMap == nil {
		return []*uint{}
	}
	var newList []*uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes uint64 as argument and returns (uint64)
// 	3. Slice of type []*uint64
//
// Returns:
//	New List of type []*uint64.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *uint64, list []*uint64) []*uint64 {
	if fFilter == nil || fMap == nil {
		return []*uint64{}
	}
	var newList []*uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes uint32 as argument and returns (uint32)
// 	3. Slice of type []*uint32
//
// Returns:
//	New List of type []*uint32.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *uint32, list []*uint32) []*uint32 {
	if fFilter == nil || fMap == nil {
		return []*uint32{}
	}
	var newList []*uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes uint16 as argument and returns (uint16)
// 	3. Slice of type []*uint16
//
// Returns:
//	New List of type []*uint16.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *uint16, list []*uint16) []*uint16 {
	if fFilter == nil || fMap == nil {
		return []*uint16{}
	}
	var newList []*uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes uint8 as argument and returns (uint8)
// 	3. Slice of type []*uint8
//
// Returns:
//	New List of type []*uint8.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *uint8, list []*uint8) []*uint8 {
	if fFilter == nil || fMap == nil {
		return []*uint8{}
	}
	var newList []*uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapStrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes string as argument and returns (string)
// 	3. Slice of type []*string
//
// Returns:
//	New List of type []*string.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrPtr(fFilter func(*string) bool, fMap func(*string) *string, list []*string) []*string {
	if fFilter == nil || fMap == nil {
		return []*string{}
	}
	var newList []*string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapBoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes bool as argument and returns (bool)
// 	3. Slice of type []*bool
//
// Returns:
//	New List of type []*bool.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolPtr(fFilter func(*bool) bool, fMap func(*bool) *bool, list []*bool) []*bool {
	if fFilter == nil || fMap == nil {
		return []*bool{}
	}
	var newList []*bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapFloat32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes float32 as argument and returns (float32)
// 	3. Slice of type []*float32
//
// Returns:
//	New List of type []*float32.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Ptr(fFilter func(*float32) bool, fMap func(*float32) *float32, list []*float32) []*float32 {
	if fFilter == nil || fMap == nil {
		return []*float32{}
	}
	var newList []*float32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapFloat64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns (bool).
//	2. Function: takes float64 as argument and returns (float64)
// 	3. Slice of type []*float64
//
// Returns:
//	New List of type []*float64.
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Ptr(fFilter func(*float64) bool, fMap func(*float64) *float64, list []*float64) []*float64 {
	if fFilter == nil || fMap == nil {
		return []*float64{}
	}
	var newList []*float64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}
