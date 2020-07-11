package fp

// RemoveInt removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveInt(isEven, []int{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num int) bool {
//		return num%2 == 0
//	}
func RemoveInt(f func(int) bool, list []int) []int {
	if f == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveInt64 removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveInt64(isEven, []int64{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num int64) bool {
//		return num%2 == 0
//	}
func RemoveInt64(f func(int64) bool, list []int64) []int64 {
	if f == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveInt32 removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveInt32(isEven, []int32{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num int32) bool {
//		return num%2 == 0
//	}
func RemoveInt32(f func(int32) bool, list []int32) []int32 {
	if f == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveInt16 removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveInt16(isEven, []int16{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num int16) bool {
//		return num%2 == 0
//	}
func RemoveInt16(f func(int16) bool, list []int16) []int16 {
	if f == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveInt8 removs the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveInt8(isEven, []int8{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num int8) bool {
//		return num%2 == 0
//	}
func RemoveInt8(f func(int8) bool, list []int8) []int8 {
	if f == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUint64 removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveUint64(isEven, []uint64{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num uint64) bool {
//		return num%2 == 0
//	}
func RemoveUint64(f func(uint64) bool, list []uint64) []uint64 {
	if f == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUint32 removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveUint32(isEven, []uint32{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num uint32) bool {
//		return num%2 == 0
//	}
func RemoveUint32(f func(uint32) bool, list []uint32) []uint32 {
	if f == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUint16 removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveUint16(isEven, []uint16{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num uint16) bool {
//		return num%2 == 0
//	}
func RemoveUint16(f func(uint16) bool, list []uint16) []uint16 {
	if f == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUint8 removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveUint8(isEven, []uint8{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num uint8) bool {
//		return num%2 == 0
//	}
func RemoveUint8(f func(uint8) bool, list []uint8) []uint8 {
	if f == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveUint removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveUint(isEven, []uint{1, 2, 3, 4}) // Returns: [1, 3]
//
//	func isEven(num uint) bool {
//		return num%2 == 0
//	}
func RemoveUint(f func(uint) bool, list []uint) []uint {
	if f == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveFloat64 removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveFloat64(isPositive, []float64{0, -2, 2, 40.50}) // Returns: [0, -2]
//
//	func isPositive(num float64) bool {
//		return num > 0
//	}
func RemoveFloat64(f func(float64) bool, list []float64) []float64 {
	if f == nil {
		return []float64{}
	}
	var newList []float64
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveFloat32 removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveFloat32(isPositive, []float32{0, -2, 2, 40.50}) // Returns: [0, -2]
//
//	func isPositive(num float32) bool {
//		return num > 0
//	}
func RemoveFloat32(f func(float32) bool, list []float32) []float32 {
	if f == nil {
		return []float32{}
	}
	var newList []float32
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveStr removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
//
// Example:
//	fp.RemoveStr(isStringLengthLessThan10, []string{"gopal", "govinda", "Nandeshwar"}) // Returns: ["Nandeshwar"]
//
//	func isStringLengthLessThan10(str string) bool {
//		return len(str) < 10
//	}
func RemoveStr(f func(string) bool, list []string) []string {
	if f == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// RemoveBool removes the items from the given list based on supplied function and returns new list
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List
//	Empty list if both of arguments are nil or either one is nil.
func RemoveBool(f func(bool) bool, list []bool) []bool {
	if f == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
