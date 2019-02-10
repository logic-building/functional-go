package fp 

// FilterMapIntInt64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt64(fFilter func(int) bool, fMap func(int) int64, list []int) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntInt32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt32(fFilter func(int) bool, fMap func(int) int32, list []int) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntInt16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt16(fFilter func(int) bool, fMap func(int) int16, list []int) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntInt8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntInt8(fFilter func(int) bool, fMap func(int) int8, list []int) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntUint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint(fFilter func(int) bool, fMap func(int) uint, list []int) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntUint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint64(fFilter func(int) bool, fMap func(int) uint64, list []int) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntUint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint32(fFilter func(int) bool, fMap func(int) uint32, list []int) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntUint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint16(fFilter func(int) bool, fMap func(int) uint16, list []int) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntUint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntUint8(fFilter func(int) bool, fMap func(int) uint8, list []int) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntStr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntStr(fFilter func(int) bool, fMap func(int) string, list []int) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapIntBool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns <OUTPUT>
// 	3. List of type int
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntBool(fFilter func(int) bool, fMap func(int) bool, list []int) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Int filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int(fFilter func(int64) bool, fMap func(int64) int, list []int64) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Int32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int32(fFilter func(int64) bool, fMap func(int64) int32, list []int64) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Int16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int16(fFilter func(int64) bool, fMap func(int64) int16, list []int64) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Int8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Int8(fFilter func(int64) bool, fMap func(int64) int8, list []int64) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Uint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint(fFilter func(int64) bool, fMap func(int64) uint, list []int64) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Uint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint64(fFilter func(int64) bool, fMap func(int64) uint64, list []int64) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Uint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint32(fFilter func(int64) bool, fMap func(int64) uint32, list []int64) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Uint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint16(fFilter func(int64) bool, fMap func(int64) uint16, list []int64) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Uint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Uint8(fFilter func(int64) bool, fMap func(int64) uint8, list []int64) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Str filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Str(fFilter func(int64) bool, fMap func(int64) string, list []int64) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt64Bool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int64 and returns true/false.
//	2. Function: takes int64 as argument and returns <OUTPUT>
// 	3. List of type int64
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt64Bool(fFilter func(int64) bool, fMap func(int64) bool, list []int64) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Int filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int(fFilter func(int32) bool, fMap func(int32) int, list []int32) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Int64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int64(fFilter func(int32) bool, fMap func(int32) int64, list []int32) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Int16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int16(fFilter func(int32) bool, fMap func(int32) int16, list []int32) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Int8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Int8(fFilter func(int32) bool, fMap func(int32) int8, list []int32) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Uint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint(fFilter func(int32) bool, fMap func(int32) uint, list []int32) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Uint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint64(fFilter func(int32) bool, fMap func(int32) uint64, list []int32) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Uint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint32(fFilter func(int32) bool, fMap func(int32) uint32, list []int32) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Uint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint16(fFilter func(int32) bool, fMap func(int32) uint16, list []int32) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Uint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Uint8(fFilter func(int32) bool, fMap func(int32) uint8, list []int32) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Str filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Str(fFilter func(int32) bool, fMap func(int32) string, list []int32) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt32Bool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int32 and returns true/false.
//	2. Function: takes int32 as argument and returns <OUTPUT>
// 	3. List of type int32
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt32Bool(fFilter func(int32) bool, fMap func(int32) bool, list []int32) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Int filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int(fFilter func(int16) bool, fMap func(int16) int, list []int16) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Int64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int64(fFilter func(int16) bool, fMap func(int16) int64, list []int16) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Int32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int32(fFilter func(int16) bool, fMap func(int16) int32, list []int16) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Int8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Int8(fFilter func(int16) bool, fMap func(int16) int8, list []int16) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Uint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint(fFilter func(int16) bool, fMap func(int16) uint, list []int16) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Uint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint64(fFilter func(int16) bool, fMap func(int16) uint64, list []int16) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Uint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint32(fFilter func(int16) bool, fMap func(int16) uint32, list []int16) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Uint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint16(fFilter func(int16) bool, fMap func(int16) uint16, list []int16) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Uint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Uint8(fFilter func(int16) bool, fMap func(int16) uint8, list []int16) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Str filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Str(fFilter func(int16) bool, fMap func(int16) string, list []int16) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt16Bool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int16 and returns true/false.
//	2. Function: takes int16 as argument and returns <OUTPUT>
// 	3. List of type int16
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt16Bool(fFilter func(int16) bool, fMap func(int16) bool, list []int16) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Int filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int(fFilter func(int8) bool, fMap func(int8) int, list []int8) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Int64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int64(fFilter func(int8) bool, fMap func(int8) int64, list []int8) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Int32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int32(fFilter func(int8) bool, fMap func(int8) int32, list []int8) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Int16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Int16(fFilter func(int8) bool, fMap func(int8) int16, list []int8) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Uint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint(fFilter func(int8) bool, fMap func(int8) uint, list []int8) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Uint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint64(fFilter func(int8) bool, fMap func(int8) uint64, list []int8) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Uint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint32(fFilter func(int8) bool, fMap func(int8) uint32, list []int8) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Uint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint16(fFilter func(int8) bool, fMap func(int8) uint16, list []int8) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Uint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Uint8(fFilter func(int8) bool, fMap func(int8) uint8, list []int8) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Str filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Str(fFilter func(int8) bool, fMap func(int8) string, list []int8) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapInt8Bool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int8 and returns true/false.
//	2. Function: takes int8 as argument and returns <OUTPUT>
// 	3. List of type int8
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInt8Bool(fFilter func(int8) bool, fMap func(int8) bool, list []int8) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintInt filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt(fFilter func(uint) bool, fMap func(uint) int, list []uint) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintInt64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt64(fFilter func(uint) bool, fMap func(uint) int64, list []uint) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintInt32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt32(fFilter func(uint) bool, fMap func(uint) int32, list []uint) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintInt16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt16(fFilter func(uint) bool, fMap func(uint) int16, list []uint) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintInt8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintInt8(fFilter func(uint) bool, fMap func(uint) int8, list []uint) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintUint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint64(fFilter func(uint) bool, fMap func(uint) uint64, list []uint) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintUint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint32(fFilter func(uint) bool, fMap func(uint) uint32, list []uint) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintUint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint16(fFilter func(uint) bool, fMap func(uint) uint16, list []uint) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintUint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintUint8(fFilter func(uint) bool, fMap func(uint) uint8, list []uint) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintStr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintStr(fFilter func(uint) bool, fMap func(uint) string, list []uint) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUintBool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint and returns true/false.
//	2. Function: takes uint as argument and returns <OUTPUT>
// 	3. List of type uint
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUintBool(fFilter func(uint) bool, fMap func(uint) bool, list []uint) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Int filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int(fFilter func(uint64) bool, fMap func(uint64) int, list []uint64) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Int64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int64(fFilter func(uint64) bool, fMap func(uint64) int64, list []uint64) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Int32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int32(fFilter func(uint64) bool, fMap func(uint64) int32, list []uint64) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Int16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int16(fFilter func(uint64) bool, fMap func(uint64) int16, list []uint64) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Int8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Int8(fFilter func(uint64) bool, fMap func(uint64) int8, list []uint64) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Uint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint(fFilter func(uint64) bool, fMap func(uint64) uint, list []uint64) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Uint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint32(fFilter func(uint64) bool, fMap func(uint64) uint32, list []uint64) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Uint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint16(fFilter func(uint64) bool, fMap func(uint64) uint16, list []uint64) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Uint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Uint8(fFilter func(uint64) bool, fMap func(uint64) uint8, list []uint64) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Str filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Str(fFilter func(uint64) bool, fMap func(uint64) string, list []uint64) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint64Bool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint64 and returns true/false.
//	2. Function: takes uint64 as argument and returns <OUTPUT>
// 	3. List of type uint64
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint64Bool(fFilter func(uint64) bool, fMap func(uint64) bool, list []uint64) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Int filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int(fFilter func(uint32) bool, fMap func(uint32) int, list []uint32) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Int64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int64(fFilter func(uint32) bool, fMap func(uint32) int64, list []uint32) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Int32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int32(fFilter func(uint32) bool, fMap func(uint32) int32, list []uint32) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Int16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int16(fFilter func(uint32) bool, fMap func(uint32) int16, list []uint32) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Int8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Int8(fFilter func(uint32) bool, fMap func(uint32) int8, list []uint32) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Uint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint(fFilter func(uint32) bool, fMap func(uint32) uint, list []uint32) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Uint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint64(fFilter func(uint32) bool, fMap func(uint32) uint64, list []uint32) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Uint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint16(fFilter func(uint32) bool, fMap func(uint32) uint16, list []uint32) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Uint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Uint8(fFilter func(uint32) bool, fMap func(uint32) uint8, list []uint32) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Str filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Str(fFilter func(uint32) bool, fMap func(uint32) string, list []uint32) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint32Bool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint32 and returns true/false.
//	2. Function: takes uint32 as argument and returns <OUTPUT>
// 	3. List of type uint32
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint32Bool(fFilter func(uint32) bool, fMap func(uint32) bool, list []uint32) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Int filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int(fFilter func(uint16) bool, fMap func(uint16) int, list []uint16) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Int64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int64(fFilter func(uint16) bool, fMap func(uint16) int64, list []uint16) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Int32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int32(fFilter func(uint16) bool, fMap func(uint16) int32, list []uint16) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Int16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int16(fFilter func(uint16) bool, fMap func(uint16) int16, list []uint16) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Int8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Int8(fFilter func(uint16) bool, fMap func(uint16) int8, list []uint16) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Uint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint(fFilter func(uint16) bool, fMap func(uint16) uint, list []uint16) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Uint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint64(fFilter func(uint16) bool, fMap func(uint16) uint64, list []uint16) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Uint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint32(fFilter func(uint16) bool, fMap func(uint16) uint32, list []uint16) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Uint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Uint8(fFilter func(uint16) bool, fMap func(uint16) uint8, list []uint16) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Str filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Str(fFilter func(uint16) bool, fMap func(uint16) string, list []uint16) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint16Bool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint16 and returns true/false.
//	2. Function: takes uint16 as argument and returns <OUTPUT>
// 	3. List of type uint16
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint16Bool(fFilter func(uint16) bool, fMap func(uint16) bool, list []uint16) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Int filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int(fFilter func(uint8) bool, fMap func(uint8) int, list []uint8) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Int64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int64(fFilter func(uint8) bool, fMap func(uint8) int64, list []uint8) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Int32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int32(fFilter func(uint8) bool, fMap func(uint8) int32, list []uint8) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Int16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int16(fFilter func(uint8) bool, fMap func(uint8) int16, list []uint8) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Int8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Int8(fFilter func(uint8) bool, fMap func(uint8) int8, list []uint8) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Uint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint(fFilter func(uint8) bool, fMap func(uint8) uint, list []uint8) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Uint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint64(fFilter func(uint8) bool, fMap func(uint8) uint64, list []uint8) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Uint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint32(fFilter func(uint8) bool, fMap func(uint8) uint32, list []uint8) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Uint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Uint16(fFilter func(uint8) bool, fMap func(uint8) uint16, list []uint8) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Str filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Str(fFilter func(uint8) bool, fMap func(uint8) string, list []uint8) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapUint8Bool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - uint8 and returns true/false.
//	2. Function: takes uint8 as argument and returns <OUTPUT>
// 	3. List of type uint8
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapUint8Bool(fFilter func(uint8) bool, fMap func(uint8) bool, list []uint8) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrInt filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt(fFilter func(string) bool, fMap func(string) int, list []string) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrInt64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt64(fFilter func(string) bool, fMap func(string) int64, list []string) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrInt32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt32(fFilter func(string) bool, fMap func(string) int32, list []string) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrInt16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt16(fFilter func(string) bool, fMap func(string) int16, list []string) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrInt8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInt8(fFilter func(string) bool, fMap func(string) int8, list []string) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrUint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint(fFilter func(string) bool, fMap func(string) uint, list []string) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrUint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint64(fFilter func(string) bool, fMap func(string) uint64, list []string) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrUint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint32(fFilter func(string) bool, fMap func(string) uint32, list []string) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrUint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint16(fFilter func(string) bool, fMap func(string) uint16, list []string) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrUint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrUint8(fFilter func(string) bool, fMap func(string) uint8, list []string) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapStrBool filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns <OUTPUT>
// 	3. List of type string
//
// Returns:
//	New List of type bool
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrBool(fFilter func(string) bool, fMap func(string) bool, list []string) []bool {
	if fFilter == nil || fMap == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolInt filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt(fFilter func(bool) bool, fMap func(bool) int, list []bool) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolInt64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type int64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt64(fFilter func(bool) bool, fMap func(bool) int64, list []bool) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolInt32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type int32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt32(fFilter func(bool) bool, fMap func(bool) int32, list []bool) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolInt16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type int16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt16(fFilter func(bool) bool, fMap func(bool) int16, list []bool) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolInt8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type int8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolInt8(fFilter func(bool) bool, fMap func(bool) int8, list []bool) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolUint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type uint
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint(fFilter func(bool) bool, fMap func(bool) uint, list []bool) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolUint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type uint64
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint64(fFilter func(bool) bool, fMap func(bool) uint64, list []bool) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolUint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type uint32
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint32(fFilter func(bool) bool, fMap func(bool) uint32, list []bool) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolUint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type uint16
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint16(fFilter func(bool) bool, fMap func(bool) uint16, list []bool) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolUint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type uint8
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolUint8(fFilter func(bool) bool, fMap func(bool) uint8, list []bool) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}


// FilterMapBoolStr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - bool and returns true/false.
//	2. Function: takes bool as argument and returns <OUTPUT>
// 	3. List of type bool
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapBoolStr(fFilter func(bool) bool, fMap func(bool) string, list []bool) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

