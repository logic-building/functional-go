package fp

// TakeWhileInt returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileInt(isEven, []int{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num int) bool {
//		return num%2 == 0
//	}
func TakeWhileInt(f func(int) bool, list []int) []int {
	if f == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileInt64 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileInt64(isEven, []int64{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num int64) bool {
//		return num%2 == 0
//	}
func TakeWhileInt64(f func(int64) bool, list []int64) []int64 {
	if f == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileInt32 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileInt32(isEven, []int32{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num int32) bool {
//		return num%2 == 0
//	}
func TakeWhileInt32(f func(int32) bool, list []int32) []int32 {
	if f == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileInt16 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileInt16(isEven, []int16{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num int16) bool {
//		return num%2 == 0
//	}
func TakeWhileInt16(f func(int16) bool, list []int16) []int16 {
	if f == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileInt8 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileInt8(isEven, []int8{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num int8) bool {
//		return num%2 == 0
//	}
func TakeWhileInt8(f func(int8) bool, list []int8) []int8 {
	if f == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUint returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileUint(isEven, []uint{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num uint) bool {
//		return num%2 == 0
//	}
func TakeWhileUint(f func(uint) bool, list []uint) []uint {
	if f == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUint64 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileUint64(isEven, []uint64{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num uint64) bool {
//		return num%2 == 0
//	}
func TakeWhileUint64(f func(uint64) bool, list []uint64) []uint64 {
	if f == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUint32 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileUint32(isEven, []uint32{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num uint32) bool {
//		return num%2 == 0
//	}
func TakeWhileUint32(f func(uint32) bool, list []uint32) []uint32 {
	if f == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUint16 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileUint16(isEven, []uint16{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num uint16) bool {
//		return num%2 == 0
//	}
func TakeWhileUint16(f func(uint16) bool, list []uint16) []uint16 {
	if f == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileUint8 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileUint8(isEven, []uint8{4, 2, 4, 7, 5}) // Returns: [4, 2, 4]
//
//	func isEven(num uint8) bool {
//		return num%2 == 0
//	}
func TakeWhileUint8(f func(uint8) bool, list []uint8) []uint8 {
	if f == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileFloat64 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileFloat64(isPositive, []float64{4, 2, 3, -1, 4, 5}) // Returns: [4, 2, 3]
//
//	func isPositive(num float64) bool {
//		return num > 0
//	}
func TakeWhileFloat64(f func(float64) bool, list []float64) []float64 {
	if f == nil {
		return []float64{}
	}
	var newList []float64
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileFloat32 returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileFloat32(isPositive, []float32{4, 2, 3, -1, 4, 5}) // Returns: [4, 2, 3]
//
//	func isPositive(num float32) bool {
//		return num > 0
//	}
func TakeWhileFloat32(f func(float32) bool, list []float32) []float32 {
	if f == nil {
		return []float32{}
	}
	var newList []float32
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileStr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
//
// Example:
//	fp.TakeWhileStr(isStringLengthLessThan10, []string{"Ram", "Shyam", "Radha", "Nandeshwar", "ShyamSundar", "Hari Shankar"}) // Returns: ["Ram", "Shyam", "Radha"]
//
//	func isStringLengthLessThan10(str string) bool {
//		return len(str) < 10
//	}
func TakeWhileStr(f func(string) bool, list []string) []string {
	if f == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

// TakeWhileBool returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function
//	2. List
//
// Returns:
//	New List.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileBool(f func(bool) bool, list []bool) []bool {
	if f == nil {
		return []bool{}
	}
	var newList []bool
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}
