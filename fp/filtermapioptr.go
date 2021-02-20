package fp

// FilterMapIntInt64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *int64
// 	3. List of type *[]int
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt64Ptr(fFilter func(*int) bool, fMap func(*int) *int64, list []*int) []*int64 {
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

// FilterMapIntInt32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *int32
// 	3. List of type *[]int
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt32Ptr(fFilter func(*int) bool, fMap func(*int) *int32, list []*int) []*int32 {
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

// FilterMapIntInt16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *int16
// 	3. List of type *[]int
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt16Ptr(fFilter func(*int) bool, fMap func(*int) *int16, list []*int) []*int16 {
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

// FilterMapIntInt8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *int8
// 	3. List of type *[]int
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt8Ptr(fFilter func(*int) bool, fMap func(*int) *int8, list []*int) []*int8 {
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

// FilterMapIntUintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *uint
// 	3. List of type *[]int
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUintPtr(fFilter func(*int) bool, fMap func(*int) *uint, list []*int) []*uint {
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

// FilterMapIntUint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *uint64
// 	3. List of type *[]int
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint64Ptr(fFilter func(*int) bool, fMap func(*int) *uint64, list []*int) []*uint64 {
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

// FilterMapIntUint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *uint32
// 	3. List of type *[]int
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint32Ptr(fFilter func(*int) bool, fMap func(*int) *uint32, list []*int) []*uint32 {
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

// FilterMapIntUint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *uint16
// 	3. List of type *[]int
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint16Ptr(fFilter func(*int) bool, fMap func(*int) *uint16, list []*int) []*uint16 {
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

// FilterMapIntUint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *uint8
// 	3. List of type *[]int
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint8Ptr(fFilter func(*int) bool, fMap func(*int) *uint8, list []*int) []*uint8 {
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

// FilterMapIntStrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *string
// 	3. List of type *[]int
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntStrPtr(fFilter func(*int) bool, fMap func(*int) *string, list []*int) []*string {
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

// FilterMapIntBoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *bool
// 	3. List of type *[]int
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntBoolPtr(fFilter func(*int) bool, fMap func(*int) *bool, list []*int) []*bool {
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

// FilterMapIntFloat32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *float32
// 	3. List of type *[]int
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntFloat32Ptr(fFilter func(*int) bool, fMap func(*int) *float32, list []*int) []*float32 {
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

// FilterMapIntFloat64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int and returns true/false.
//	2. Function: takes *int as argument and returns *float64
// 	3. List of type *[]int
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntFloat64Ptr(fFilter func(*int) bool, fMap func(*int) *float64, list []*int) []*float64 {
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

// FilterMapInt64IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *int
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64IntPtr(fFilter func(*int64) bool, fMap func(*int64) *int, list []*int64) []*int {
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

// FilterMapInt64Int32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *int32
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int32Ptr(fFilter func(*int64) bool, fMap func(*int64) *int32, list []*int64) []*int32 {
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

// FilterMapInt64Int16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *int16
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int16Ptr(fFilter func(*int64) bool, fMap func(*int64) *int16, list []*int64) []*int16 {
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

// FilterMapInt64Int8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *int8
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int8Ptr(fFilter func(*int64) bool, fMap func(*int64) *int8, list []*int64) []*int8 {
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

// FilterMapInt64UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *uint
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64UintPtr(fFilter func(*int64) bool, fMap func(*int64) *uint, list []*int64) []*uint {
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

// FilterMapInt64Uint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *uint64
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint64Ptr(fFilter func(*int64) bool, fMap func(*int64) *uint64, list []*int64) []*uint64 {
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

// FilterMapInt64Uint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *uint32
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint32Ptr(fFilter func(*int64) bool, fMap func(*int64) *uint32, list []*int64) []*uint32 {
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

// FilterMapInt64Uint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *uint16
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint16Ptr(fFilter func(*int64) bool, fMap func(*int64) *uint16, list []*int64) []*uint16 {
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

// FilterMapInt64Uint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *uint8
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint8Ptr(fFilter func(*int64) bool, fMap func(*int64) *uint8, list []*int64) []*uint8 {
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

// FilterMapInt64StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *string
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64StrPtr(fFilter func(*int64) bool, fMap func(*int64) *string, list []*int64) []*string {
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

// FilterMapInt64BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *bool
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64BoolPtr(fFilter func(*int64) bool, fMap func(*int64) *bool, list []*int64) []*bool {
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

// FilterMapInt64Float32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *float32
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Float32Ptr(fFilter func(*int64) bool, fMap func(*int64) *float32, list []*int64) []*float32 {
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

// FilterMapInt64Float64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int64 and returns true/false.
//	2. Function: takes *int64 as argument and returns *float64
// 	3. List of type *[]int64
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Float64Ptr(fFilter func(*int64) bool, fMap func(*int64) *float64, list []*int64) []*float64 {
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

// FilterMapInt32IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *int
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32IntPtr(fFilter func(*int32) bool, fMap func(*int32) *int, list []*int32) []*int {
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

// FilterMapInt32Int64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *int64
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int64Ptr(fFilter func(*int32) bool, fMap func(*int32) *int64, list []*int32) []*int64 {
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

// FilterMapInt32Int16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *int16
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int16Ptr(fFilter func(*int32) bool, fMap func(*int32) *int16, list []*int32) []*int16 {
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

// FilterMapInt32Int8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *int8
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int8Ptr(fFilter func(*int32) bool, fMap func(*int32) *int8, list []*int32) []*int8 {
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

// FilterMapInt32UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *uint
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32UintPtr(fFilter func(*int32) bool, fMap func(*int32) *uint, list []*int32) []*uint {
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

// FilterMapInt32Uint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *uint64
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint64Ptr(fFilter func(*int32) bool, fMap func(*int32) *uint64, list []*int32) []*uint64 {
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

// FilterMapInt32Uint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *uint32
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint32Ptr(fFilter func(*int32) bool, fMap func(*int32) *uint32, list []*int32) []*uint32 {
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

// FilterMapInt32Uint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *uint16
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint16Ptr(fFilter func(*int32) bool, fMap func(*int32) *uint16, list []*int32) []*uint16 {
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

// FilterMapInt32Uint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *uint8
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint8Ptr(fFilter func(*int32) bool, fMap func(*int32) *uint8, list []*int32) []*uint8 {
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

// FilterMapInt32StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *string
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32StrPtr(fFilter func(*int32) bool, fMap func(*int32) *string, list []*int32) []*string {
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

// FilterMapInt32BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *bool
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32BoolPtr(fFilter func(*int32) bool, fMap func(*int32) *bool, list []*int32) []*bool {
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

// FilterMapInt32Float32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *float32
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Float32Ptr(fFilter func(*int32) bool, fMap func(*int32) *float32, list []*int32) []*float32 {
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

// FilterMapInt32Float64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int32 and returns true/false.
//	2. Function: takes *int32 as argument and returns *float64
// 	3. List of type *[]int32
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Float64Ptr(fFilter func(*int32) bool, fMap func(*int32) *float64, list []*int32) []*float64 {
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

// FilterMapInt16IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *int
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16IntPtr(fFilter func(*int16) bool, fMap func(*int16) *int, list []*int16) []*int {
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

// FilterMapInt16Int64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *int64
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int64Ptr(fFilter func(*int16) bool, fMap func(*int16) *int64, list []*int16) []*int64 {
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

// FilterMapInt16Int32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *int32
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int32Ptr(fFilter func(*int16) bool, fMap func(*int16) *int32, list []*int16) []*int32 {
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

// FilterMapInt16Int8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *int8
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int8Ptr(fFilter func(*int16) bool, fMap func(*int16) *int8, list []*int16) []*int8 {
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

// FilterMapInt16UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *uint
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16UintPtr(fFilter func(*int16) bool, fMap func(*int16) *uint, list []*int16) []*uint {
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

// FilterMapInt16Uint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *uint64
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint64Ptr(fFilter func(*int16) bool, fMap func(*int16) *uint64, list []*int16) []*uint64 {
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

// FilterMapInt16Uint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *uint32
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint32Ptr(fFilter func(*int16) bool, fMap func(*int16) *uint32, list []*int16) []*uint32 {
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

// FilterMapInt16Uint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *uint16
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint16Ptr(fFilter func(*int16) bool, fMap func(*int16) *uint16, list []*int16) []*uint16 {
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

// FilterMapInt16Uint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *uint8
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint8Ptr(fFilter func(*int16) bool, fMap func(*int16) *uint8, list []*int16) []*uint8 {
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

// FilterMapInt16StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *string
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16StrPtr(fFilter func(*int16) bool, fMap func(*int16) *string, list []*int16) []*string {
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

// FilterMapInt16BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *bool
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16BoolPtr(fFilter func(*int16) bool, fMap func(*int16) *bool, list []*int16) []*bool {
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

// FilterMapInt16Float32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *float32
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Float32Ptr(fFilter func(*int16) bool, fMap func(*int16) *float32, list []*int16) []*float32 {
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

// FilterMapInt16Float64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int16 and returns true/false.
//	2. Function: takes *int16 as argument and returns *float64
// 	3. List of type *[]int16
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Float64Ptr(fFilter func(*int16) bool, fMap func(*int16) *float64, list []*int16) []*float64 {
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

// FilterMapInt8IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *int
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8IntPtr(fFilter func(*int8) bool, fMap func(*int8) *int, list []*int8) []*int {
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

// FilterMapInt8Int64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *int64
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int64Ptr(fFilter func(*int8) bool, fMap func(*int8) *int64, list []*int8) []*int64 {
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

// FilterMapInt8Int32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *int32
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int32Ptr(fFilter func(*int8) bool, fMap func(*int8) *int32, list []*int8) []*int32 {
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

// FilterMapInt8Int16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *int16
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int16Ptr(fFilter func(*int8) bool, fMap func(*int8) *int16, list []*int8) []*int16 {
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

// FilterMapInt8UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *uint
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8UintPtr(fFilter func(*int8) bool, fMap func(*int8) *uint, list []*int8) []*uint {
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

// FilterMapInt8Uint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *uint64
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint64Ptr(fFilter func(*int8) bool, fMap func(*int8) *uint64, list []*int8) []*uint64 {
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

// FilterMapInt8Uint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *uint32
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint32Ptr(fFilter func(*int8) bool, fMap func(*int8) *uint32, list []*int8) []*uint32 {
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

// FilterMapInt8Uint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *uint16
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint16Ptr(fFilter func(*int8) bool, fMap func(*int8) *uint16, list []*int8) []*uint16 {
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

// FilterMapInt8Uint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *uint8
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint8Ptr(fFilter func(*int8) bool, fMap func(*int8) *uint8, list []*int8) []*uint8 {
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

// FilterMapInt8StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *string
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8StrPtr(fFilter func(*int8) bool, fMap func(*int8) *string, list []*int8) []*string {
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

// FilterMapInt8BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *bool
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8BoolPtr(fFilter func(*int8) bool, fMap func(*int8) *bool, list []*int8) []*bool {
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

// FilterMapInt8Float32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *float32
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Float32Ptr(fFilter func(*int8) bool, fMap func(*int8) *float32, list []*int8) []*float32 {
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

// FilterMapInt8Float64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *int8 and returns true/false.
//	2. Function: takes *int8 as argument and returns *float64
// 	3. List of type *[]int8
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Float64Ptr(fFilter func(*int8) bool, fMap func(*int8) *float64, list []*int8) []*float64 {
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

// FilterMapUintIntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *int
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintIntPtr(fFilter func(*uint) bool, fMap func(*uint) *int, list []*uint) []*int {
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

// FilterMapUintInt64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *int64
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt64Ptr(fFilter func(*uint) bool, fMap func(*uint) *int64, list []*uint) []*int64 {
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

// FilterMapUintInt32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *int32
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt32Ptr(fFilter func(*uint) bool, fMap func(*uint) *int32, list []*uint) []*int32 {
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

// FilterMapUintInt16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *int16
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt16Ptr(fFilter func(*uint) bool, fMap func(*uint) *int16, list []*uint) []*int16 {
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

// FilterMapUintInt8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *int8
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt8Ptr(fFilter func(*uint) bool, fMap func(*uint) *int8, list []*uint) []*int8 {
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

// FilterMapUintUint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *uint64
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint64Ptr(fFilter func(*uint) bool, fMap func(*uint) *uint64, list []*uint) []*uint64 {
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

// FilterMapUintUint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *uint32
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint32Ptr(fFilter func(*uint) bool, fMap func(*uint) *uint32, list []*uint) []*uint32 {
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

// FilterMapUintUint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *uint16
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint16Ptr(fFilter func(*uint) bool, fMap func(*uint) *uint16, list []*uint) []*uint16 {
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

// FilterMapUintUint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *uint8
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint8Ptr(fFilter func(*uint) bool, fMap func(*uint) *uint8, list []*uint) []*uint8 {
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

// FilterMapUintStrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *string
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintStrPtr(fFilter func(*uint) bool, fMap func(*uint) *string, list []*uint) []*string {
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

// FilterMapUintBoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *bool
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintBoolPtr(fFilter func(*uint) bool, fMap func(*uint) *bool, list []*uint) []*bool {
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

// FilterMapUintFloat32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *float32
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintFloat32Ptr(fFilter func(*uint) bool, fMap func(*uint) *float32, list []*uint) []*float32 {
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

// FilterMapUintFloat64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint and returns true/false.
//	2. Function: takes *uint as argument and returns *float64
// 	3. List of type *[]uint
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintFloat64Ptr(fFilter func(*uint) bool, fMap func(*uint) *float64, list []*uint) []*float64 {
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

// FilterMapUint64IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *int
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64IntPtr(fFilter func(*uint64) bool, fMap func(*uint64) *int, list []*uint64) []*int {
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

// FilterMapUint64Int64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *int64
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int64Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *int64, list []*uint64) []*int64 {
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

// FilterMapUint64Int32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *int32
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int32Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *int32, list []*uint64) []*int32 {
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

// FilterMapUint64Int16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *int16
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int16Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *int16, list []*uint64) []*int16 {
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

// FilterMapUint64Int8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *int8
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int8Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *int8, list []*uint64) []*int8 {
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

// FilterMapUint64UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *uint
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64UintPtr(fFilter func(*uint64) bool, fMap func(*uint64) *uint, list []*uint64) []*uint {
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

// FilterMapUint64Uint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *uint32
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint32Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *uint32, list []*uint64) []*uint32 {
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

// FilterMapUint64Uint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *uint16
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint16Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *uint16, list []*uint64) []*uint16 {
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

// FilterMapUint64Uint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *uint8
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint8Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *uint8, list []*uint64) []*uint8 {
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

// FilterMapUint64StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *string
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64StrPtr(fFilter func(*uint64) bool, fMap func(*uint64) *string, list []*uint64) []*string {
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

// FilterMapUint64BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *bool
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64BoolPtr(fFilter func(*uint64) bool, fMap func(*uint64) *bool, list []*uint64) []*bool {
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

// FilterMapUint64Float32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *float32
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Float32Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *float32, list []*uint64) []*float32 {
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

// FilterMapUint64Float64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint64 and returns true/false.
//	2. Function: takes *uint64 as argument and returns *float64
// 	3. List of type *[]uint64
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Float64Ptr(fFilter func(*uint64) bool, fMap func(*uint64) *float64, list []*uint64) []*float64 {
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

// FilterMapUint32IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *int
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32IntPtr(fFilter func(*uint32) bool, fMap func(*uint32) *int, list []*uint32) []*int {
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

// FilterMapUint32Int64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *int64
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int64Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *int64, list []*uint32) []*int64 {
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

// FilterMapUint32Int32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *int32
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int32Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *int32, list []*uint32) []*int32 {
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

// FilterMapUint32Int16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *int16
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int16Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *int16, list []*uint32) []*int16 {
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

// FilterMapUint32Int8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *int8
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int8Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *int8, list []*uint32) []*int8 {
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

// FilterMapUint32UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *uint
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32UintPtr(fFilter func(*uint32) bool, fMap func(*uint32) *uint, list []*uint32) []*uint {
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

// FilterMapUint32Uint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *uint64
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint64Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *uint64, list []*uint32) []*uint64 {
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

// FilterMapUint32Uint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *uint16
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint16Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *uint16, list []*uint32) []*uint16 {
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

// FilterMapUint32Uint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *uint8
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint8Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *uint8, list []*uint32) []*uint8 {
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

// FilterMapUint32StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *string
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32StrPtr(fFilter func(*uint32) bool, fMap func(*uint32) *string, list []*uint32) []*string {
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

// FilterMapUint32BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *bool
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32BoolPtr(fFilter func(*uint32) bool, fMap func(*uint32) *bool, list []*uint32) []*bool {
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

// FilterMapUint32Float32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *float32
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Float32Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *float32, list []*uint32) []*float32 {
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

// FilterMapUint32Float64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint32 and returns true/false.
//	2. Function: takes *uint32 as argument and returns *float64
// 	3. List of type *[]uint32
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Float64Ptr(fFilter func(*uint32) bool, fMap func(*uint32) *float64, list []*uint32) []*float64 {
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

// FilterMapUint16IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *int
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16IntPtr(fFilter func(*uint16) bool, fMap func(*uint16) *int, list []*uint16) []*int {
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

// FilterMapUint16Int64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *int64
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int64Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *int64, list []*uint16) []*int64 {
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

// FilterMapUint16Int32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *int32
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int32Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *int32, list []*uint16) []*int32 {
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

// FilterMapUint16Int16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *int16
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int16Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *int16, list []*uint16) []*int16 {
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

// FilterMapUint16Int8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *int8
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int8Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *int8, list []*uint16) []*int8 {
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

// FilterMapUint16UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *uint
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16UintPtr(fFilter func(*uint16) bool, fMap func(*uint16) *uint, list []*uint16) []*uint {
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

// FilterMapUint16Uint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *uint64
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint64Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *uint64, list []*uint16) []*uint64 {
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

// FilterMapUint16Uint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *uint32
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint32Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *uint32, list []*uint16) []*uint32 {
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

// FilterMapUint16Uint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *uint8
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint8Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *uint8, list []*uint16) []*uint8 {
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

// FilterMapUint16StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *string
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16StrPtr(fFilter func(*uint16) bool, fMap func(*uint16) *string, list []*uint16) []*string {
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

// FilterMapUint16BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *bool
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16BoolPtr(fFilter func(*uint16) bool, fMap func(*uint16) *bool, list []*uint16) []*bool {
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

// FilterMapUint16Float32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *float32
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Float32Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *float32, list []*uint16) []*float32 {
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

// FilterMapUint16Float64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint16 and returns true/false.
//	2. Function: takes *uint16 as argument and returns *float64
// 	3. List of type *[]uint16
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Float64Ptr(fFilter func(*uint16) bool, fMap func(*uint16) *float64, list []*uint16) []*float64 {
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

// FilterMapUint8IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *int
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8IntPtr(fFilter func(*uint8) bool, fMap func(*uint8) *int, list []*uint8) []*int {
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

// FilterMapUint8Int64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *int64
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int64Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *int64, list []*uint8) []*int64 {
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

// FilterMapUint8Int32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *int32
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int32Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *int32, list []*uint8) []*int32 {
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

// FilterMapUint8Int16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *int16
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int16Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *int16, list []*uint8) []*int16 {
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

// FilterMapUint8Int8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *int8
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int8Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *int8, list []*uint8) []*int8 {
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

// FilterMapUint8UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *uint
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8UintPtr(fFilter func(*uint8) bool, fMap func(*uint8) *uint, list []*uint8) []*uint {
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

// FilterMapUint8Uint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *uint64
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint64Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *uint64, list []*uint8) []*uint64 {
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

// FilterMapUint8Uint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *uint32
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint32Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *uint32, list []*uint8) []*uint32 {
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

// FilterMapUint8Uint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *uint16
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint16Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *uint16, list []*uint8) []*uint16 {
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

// FilterMapUint8StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *string
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8StrPtr(fFilter func(*uint8) bool, fMap func(*uint8) *string, list []*uint8) []*string {
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

// FilterMapUint8BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *bool
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8BoolPtr(fFilter func(*uint8) bool, fMap func(*uint8) *bool, list []*uint8) []*bool {
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

// FilterMapUint8Float32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *float32
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Float32Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *float32, list []*uint8) []*float32 {
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

// FilterMapUint8Float64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *uint8 and returns true/false.
//	2. Function: takes *uint8 as argument and returns *float64
// 	3. List of type *[]uint8
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Float64Ptr(fFilter func(*uint8) bool, fMap func(*uint8) *float64, list []*uint8) []*float64 {
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

// FilterMapStrIntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *int
// 	3. List of type *[]string
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrIntPtr(fFilter func(*string) bool, fMap func(*string) *int, list []*string) []*int {
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

// FilterMapStrInt64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *int64
// 	3. List of type *[]string
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt64Ptr(fFilter func(*string) bool, fMap func(*string) *int64, list []*string) []*int64 {
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

// FilterMapStrInt32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *int32
// 	3. List of type *[]string
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt32Ptr(fFilter func(*string) bool, fMap func(*string) *int32, list []*string) []*int32 {
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

// FilterMapStrInt16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *int16
// 	3. List of type *[]string
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt16Ptr(fFilter func(*string) bool, fMap func(*string) *int16, list []*string) []*int16 {
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

// FilterMapStrInt8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *int8
// 	3. List of type *[]string
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt8Ptr(fFilter func(*string) bool, fMap func(*string) *int8, list []*string) []*int8 {
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

// FilterMapStrUintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *uint
// 	3. List of type *[]string
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUintPtr(fFilter func(*string) bool, fMap func(*string) *uint, list []*string) []*uint {
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

// FilterMapStrUint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *uint64
// 	3. List of type *[]string
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint64Ptr(fFilter func(*string) bool, fMap func(*string) *uint64, list []*string) []*uint64 {
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

// FilterMapStrUint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *uint32
// 	3. List of type *[]string
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint32Ptr(fFilter func(*string) bool, fMap func(*string) *uint32, list []*string) []*uint32 {
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

// FilterMapStrUint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *uint16
// 	3. List of type *[]string
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint16Ptr(fFilter func(*string) bool, fMap func(*string) *uint16, list []*string) []*uint16 {
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

// FilterMapStrUint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *uint8
// 	3. List of type *[]string
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint8Ptr(fFilter func(*string) bool, fMap func(*string) *uint8, list []*string) []*uint8 {
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

// FilterMapStrBoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *bool
// 	3. List of type *[]string
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrBoolPtr(fFilter func(*string) bool, fMap func(*string) *bool, list []*string) []*bool {
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

// FilterMapStrFloat32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *float32
// 	3. List of type *[]string
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrFloat32Ptr(fFilter func(*string) bool, fMap func(*string) *float32, list []*string) []*float32 {
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

// FilterMapStrFloat64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *string and returns true/false.
//	2. Function: takes *string as argument and returns *float64
// 	3. List of type *[]string
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrFloat64Ptr(fFilter func(*string) bool, fMap func(*string) *float64, list []*string) []*float64 {
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

// FilterMapBoolIntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *int
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolIntPtr(fFilter func(*bool) bool, fMap func(*bool) *int, list []*bool) []*int {
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

// FilterMapBoolInt64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *int64
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt64Ptr(fFilter func(*bool) bool, fMap func(*bool) *int64, list []*bool) []*int64 {
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

// FilterMapBoolInt32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *int32
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt32Ptr(fFilter func(*bool) bool, fMap func(*bool) *int32, list []*bool) []*int32 {
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

// FilterMapBoolInt16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *int16
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt16Ptr(fFilter func(*bool) bool, fMap func(*bool) *int16, list []*bool) []*int16 {
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

// FilterMapBoolInt8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *int8
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt8Ptr(fFilter func(*bool) bool, fMap func(*bool) *int8, list []*bool) []*int8 {
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

// FilterMapBoolUintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *uint
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUintPtr(fFilter func(*bool) bool, fMap func(*bool) *uint, list []*bool) []*uint {
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

// FilterMapBoolUint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *uint64
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint64Ptr(fFilter func(*bool) bool, fMap func(*bool) *uint64, list []*bool) []*uint64 {
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

// FilterMapBoolUint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *uint32
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint32Ptr(fFilter func(*bool) bool, fMap func(*bool) *uint32, list []*bool) []*uint32 {
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

// FilterMapBoolUint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *uint16
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint16Ptr(fFilter func(*bool) bool, fMap func(*bool) *uint16, list []*bool) []*uint16 {
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

// FilterMapBoolUint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *uint8
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint8Ptr(fFilter func(*bool) bool, fMap func(*bool) *uint8, list []*bool) []*uint8 {
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

// FilterMapBoolStrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *string
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolStrPtr(fFilter func(*bool) bool, fMap func(*bool) *string, list []*bool) []*string {
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

// FilterMapBoolFloat32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *float32
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolFloat32Ptr(fFilter func(*bool) bool, fMap func(*bool) *float32, list []*bool) []*float32 {
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

// FilterMapBoolFloat64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *bool and returns true/false.
//	2. Function: takes *bool as argument and returns *float64
// 	3. List of type *[]bool
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolFloat64Ptr(fFilter func(*bool) bool, fMap func(*bool) *float64, list []*bool) []*float64 {
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

// FilterMapFloat32IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *int
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32IntPtr(fFilter func(*float32) bool, fMap func(*float32) *int, list []*float32) []*int {
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

// FilterMapFloat32Int64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *int64
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int64Ptr(fFilter func(*float32) bool, fMap func(*float32) *int64, list []*float32) []*int64 {
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

// FilterMapFloat32Int32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *int32
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int32Ptr(fFilter func(*float32) bool, fMap func(*float32) *int32, list []*float32) []*int32 {
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

// FilterMapFloat32Int16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *int16
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int16Ptr(fFilter func(*float32) bool, fMap func(*float32) *int16, list []*float32) []*int16 {
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

// FilterMapFloat32Int8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *int8
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Int8Ptr(fFilter func(*float32) bool, fMap func(*float32) *int8, list []*float32) []*int8 {
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

// FilterMapFloat32UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *uint
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32UintPtr(fFilter func(*float32) bool, fMap func(*float32) *uint, list []*float32) []*uint {
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

// FilterMapFloat32Uint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *uint64
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint64Ptr(fFilter func(*float32) bool, fMap func(*float32) *uint64, list []*float32) []*uint64 {
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

// FilterMapFloat32Uint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *uint32
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint32Ptr(fFilter func(*float32) bool, fMap func(*float32) *uint32, list []*float32) []*uint32 {
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

// FilterMapFloat32Uint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *uint16
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint16Ptr(fFilter func(*float32) bool, fMap func(*float32) *uint16, list []*float32) []*uint16 {
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

// FilterMapFloat32Uint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *uint8
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Uint8Ptr(fFilter func(*float32) bool, fMap func(*float32) *uint8, list []*float32) []*uint8 {
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

// FilterMapFloat32StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *string
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32StrPtr(fFilter func(*float32) bool, fMap func(*float32) *string, list []*float32) []*string {
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

// FilterMapFloat32BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *bool
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32BoolPtr(fFilter func(*float32) bool, fMap func(*float32) *bool, list []*float32) []*bool {
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

// FilterMapFloat32Float64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float32 and returns true/false.
//	2. Function: takes *float32 as argument and returns *float64
// 	3. List of type *[]float32
//
// Returns:
//	New List of type *float64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat32Float64Ptr(fFilter func(*float32) bool, fMap func(*float32) *float64, list []*float32) []*float64 {
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

// FilterMapFloat64IntPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *int
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64IntPtr(fFilter func(*float64) bool, fMap func(*float64) *int, list []*float64) []*int {
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

// FilterMapFloat64Int64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *int64
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int64Ptr(fFilter func(*float64) bool, fMap func(*float64) *int64, list []*float64) []*int64 {
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

// FilterMapFloat64Int32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *int32
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int32Ptr(fFilter func(*float64) bool, fMap func(*float64) *int32, list []*float64) []*int32 {
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

// FilterMapFloat64Int16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *int16
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int16Ptr(fFilter func(*float64) bool, fMap func(*float64) *int16, list []*float64) []*int16 {
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

// FilterMapFloat64Int8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *int8
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Int8Ptr(fFilter func(*float64) bool, fMap func(*float64) *int8, list []*float64) []*int8 {
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

// FilterMapFloat64UintPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *uint
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64UintPtr(fFilter func(*float64) bool, fMap func(*float64) *uint, list []*float64) []*uint {
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

// FilterMapFloat64Uint64Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *uint64
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint64Ptr(fFilter func(*float64) bool, fMap func(*float64) *uint64, list []*float64) []*uint64 {
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

// FilterMapFloat64Uint32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *uint32
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint32Ptr(fFilter func(*float64) bool, fMap func(*float64) *uint32, list []*float64) []*uint32 {
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

// FilterMapFloat64Uint16Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *uint16
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint16Ptr(fFilter func(*float64) bool, fMap func(*float64) *uint16, list []*float64) []*uint16 {
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

// FilterMapFloat64Uint8Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *uint8
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Uint8Ptr(fFilter func(*float64) bool, fMap func(*float64) *uint8, list []*float64) []*uint8 {
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

// FilterMapFloat64StrPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *string
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64StrPtr(fFilter func(*float64) bool, fMap func(*float64) *string, list []*float64) []*string {
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

// FilterMapFloat64BoolPtr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *bool
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64BoolPtr(fFilter func(*float64) bool, fMap func(*float64) *bool, list []*float64) []*bool {
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

// FilterMapFloat64Float32Ptr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - *float64 and returns true/false.
//	2. Function: takes *float64 as argument and returns *float32
// 	3. List of type *[]float64
//
// Returns:
//	New List of type *float32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapFloat64Float32Ptr(fFilter func(*float64) bool, fMap func(*float64) *float32, list []*float64) []*float32 {
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
