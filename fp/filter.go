package fp

// FilterInt filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all even numbers
//	fp.FilterInt(isEven, []int{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num int) bool {
//		return num%2 == 0
//	}
func FilterInt(f func(int) bool, list []int) []int {
	if f == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterInt64 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: FilterInt64 all even numbers
//	fp.FilterInt(isEven, []int64{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num int64) bool {
//		return num%2 == 0
//	}
func FilterInt64(f func(int64) bool, list []int64) []int64 {
	if f == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterInt32 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all even numbers
//	fp.FilterInt(isEven, []int32{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num int32) bool {
//		return num%2 == 0
//	}
func FilterInt32(f func(int32) bool, list []int32) []int32 {
	if f == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterInt16 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all even numbers
//	fp.FilterInt16(isEven, []int32{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num int32) bool {
//		return num%2 == 0
//	}
func FilterInt16(f func(int16) bool, list []int16) []int16 {
	if f == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterInt8 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all even numbers
//	fp.FilterInt8(isEven, []int8{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num int8) bool {
//		return num%2 == 0
//	}
func FilterInt8(f func(int8) bool, list []int8) []int8 {
	if f == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUint64 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all even numbers
//	fp.FilterUint64(isEven, []uint64{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num uint64) bool {
//		return num%2 == 0
//	}
func FilterUint64(f func(uint64) bool, list []uint64) []uint64 {
	if f == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUint32 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all even numbers
//	fp.FilterUint32(isEven, []uint32{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num uint32) bool {
//		return num%2 == 0
//	}
func FilterUint32(f func(uint32) bool, list []uint32) []uint32 {
	if f == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUint16 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all even numbers
//	fp.FilterUint16(isEven, []uint16{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num uint16) bool {
//		return num%2 == 0
//	}
func FilterUint16(f func(uint16) bool, list []uint16) []uint16 {
	if f == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUint8 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all even numbers
//	fp.FilterUint8(isEven, []uint8{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num uint8) bool {
//		return num%2 == 0
//	}
func FilterUint8(f func(uint8) bool, list []uint8) []uint8 {
	if f == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUint filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all even numbers
//	fp.FilterUint(isEven, []uint{1, 2, 3, 4}) // Returns [2, 4]
//
//	func isEven(num uint) bool {
//		return num%2 == 0
//	}
func FilterUint(f func(uint) bool, list []uint) []uint {
	if f == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterFloat64 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all positive numbers
//	fp.FilterFloat64(isPositive, []float64{0, -2, 2, 40.50}) // Returns [2, 40.50]
//
//	func isPositive(num float64) bool {
//		return num > 0
//	}
func FilterFloat64(f func(float64) bool, list []float64) []float64 {
	if f == nil {
		return []float64{}
	}
	var newList []float64
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterFloat32 filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all positive numbers
//	fp.FilterFloat32(isPositive, []float32{0, -2, 2, 40.50}) // Returns [2, 40.50]
//
//	func isPositive(num float32) bool {
//		return num > 0
//	}
func FilterFloat32(f func(float32) bool, list []float32) []float32 {
	if f == nil {
		return []float32{}
	}
	var newList []float32
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterStr filters list based on function passed as 1st argument
//
// Takes two inputs
//	1. function - Takes 1 input and returns bool
// 	2. list
//
// Returns:
//	New Filtered List.
//	Empty list if either of arguments or both of them are nil
//
// Example: Filter all the name which length is less than 10
//	fp.FilterStr(isStringLengthLessThan10, []string{"gopal", "govinda", "Dasananda Ravan"}) // Returns ["gopal", "govinda"]
//
//	func isStringLengthLessThan10(str string) bool {
//		return len(str) < 10
//	}
func FilterStr(f func(string) bool, list []string) []string {
	if f == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterBool takes two arguments
//  1. Function: takes 1 argument of type bool and returns bool
//  2. slice of type []bool
//
// Returns:
//  new filtered list
func FilterBool(f func(bool) bool, list []bool) []bool {
	if f == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
