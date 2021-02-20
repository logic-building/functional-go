package fp

// SomeInt finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeInt(isEven, []int{8, 2, 10, 4}) // Returns true
//	fp.SomeInt(isEven, []int{1, 3, 5, 7}) // Returns false
//	fp.SomeInt(nil, nil) // Returns false
//	fp.SomeInt(isEven, []int{}) // Returns false
//
//	func isEven(num int) bool {
//		return num%2 == 0
//	}
func SomeInt(f func(int) bool, list []int) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeInt64 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeInt64(isEven, []int64{8, 2, 10, 4}) // Returns true
//	fp.SomeInt64(isEven, []int64{1, 3, 5, 7}) // Returns false
//	fp.SomeInt64(nil, nil) // Returns false
//	fp.SomeInt64(isEven, []int64{}) // Returns false
//
//	func isEven(num int64) bool {
//		return num%2 == 0
//	}
func SomeInt64(f func(int64) bool, list []int64) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeInt32 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeInt32(isEven, []int32{8, 2, 10, 4}) // Returns true
//	fp.SomeInt32(isEven, []int32{1, 3, 5, 7}) // Returns false
//	fp.SomeInt32(nil, nil) // Returns false
//	fp.SomeInt32(isEven, []int32{}) // Returns false
//
//	func isEven(num int32) bool {
//		return num%2 == 0
//	}
func SomeInt32(f func(int32) bool, list []int32) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeInt16 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeInt16(isEven, []int16{8, 2, 10, 4}) // Returns true
//	fp.SomeInt16(isEven, []int16{1, 3, 5, 7}) // Returns false
//	fp.SomeInt16(nil, nil) // Returns false
//	fp.SomeInt16(isEven, []int16{}) // Returns false
//
//	func isEven(num int16) bool {
//		return num%2 == 0
//	}
func SomeInt16(f func(int16) bool, list []int16) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeInt8 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeInt8(isEven, []int16{8, 2, 10, 4}) // Returns true
//	fp.SomeInt8(isEven, []int16{1, 3, 5, 7}) // Returns false
//	fp.SomeInt8(nil, nil) // Returns false
//	fp.SomeInt8(isEven, []int16{}) // Returns false
//
//	func isEven(num int16) bool {
//		return num%2 == 0
//	}
func SomeInt8(f func(int8) bool, list []int8) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeUint64 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeUint64(isEven, []uint64{8, 2, 10, 4}) // Returns true
//	fp.SomeUint64(isEven, []uint64{1, 3, 5, 7}) // Returns false
//	fp.SomeUint64(nil, nil) // Returns false
//	fp.SomeUint64(isEven, []uint64{}) // Returns false
//
//	func isEven(num int16) bool {
//		return num%2 == 0
//	}
func SomeUint64(f func(uint64) bool, list []uint64) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeUint32 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeUint32(isEven, []uint32{8, 2, 10, 4}) // Returns true
//	fp.SomeUint32(isEven, []uint32{1, 3, 5, 7}) // Returns false
//	fp.SomeUint32(nil, nil) // Returns false
//	fp.SomeUint32(isEven, []uint32{}) // Returns false
//
//	func isEven(num uint32) bool {
//		return num%2 == 0
//	}
func SomeUint32(f func(uint32) bool, list []uint32) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeUint16 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeUint16(isEven, []uint16{8, 2, 10, 4}) // Returns true
//	fp.SomeUint16(isEven, []uint16{1, 3, 5, 7}) // Returns false
//	fp.SomeUint16(nil, nil) // Returns false
//	fp.SomeUint16(isEven, []uint16{}) // Returns false
//
//	func isEven(num uint16) bool {
//		return num%2 == 0
//	}
func SomeUint16(f func(uint16) bool, list []uint16) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeUint8 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeUint8(isEven, []uint8{8, 2, 10, 4}) // Returns true
//	fp.SomeUint8(isEven, []uint8{1, 3, 5, 7}) // Returns false
//	fp.SomeUint8(nil, nil) // Returns false
//	fp.SomeUint8(isEven, []uint8{}) // Returns false
//
//	func isEven(num uint8) bool {
//		return num%2 == 0
//	}
func SomeUint8(f func(uint8) bool, list []uint8) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeUint finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeUint(isEven, []uint{8, 2, 10, 4}) // Returns true
//	fp.SomeUint(isEven, []uint{1, 3, 5, 7}) // Returns false
//	fp.SomeUint(nil, nil) // Returns false
//	fp.SomeUint(isEven, []uint{}) // Returns false
//
//	func isEven(num uint) bool {
//		return num%2 == 0
//	}
func SomeUint(f func(uint) bool, list []uint) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeFloat64 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeFloat64(isPositive, []float64{8, 2, 10, 4}) // Returns true
//	fp.SomeFloat64(nil, nil) // Returns false
//	fp.SomeFloat64(isPositive, []float64{}) // Returns false
//
//	func isPositive(num float64) bool {
//		return num > 0
//	}
func SomeFloat64(f func(float64) bool, list []float64) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeFloat32 finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeFloat32(isPositive, []float32{8, 2, 10, 4}) // Returns true
//	fp.SomeFloat32(nil, nil) // Returns false
//	fp.SomeFloat32(isPositive, []float32{}) // Returns false
//
//	func isPositive(num float32) bool {
//		return num > 0
//	}
func SomeFloat32(f func(float32) bool, list []float32) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

// SomeStr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
//
// Example:
//	fp.SomeStr(isStringLengthLessThan10, []string{"gopal", "govinda", "Nandeshwar"}) // Returns true
//	fp.SomeStr(nil, nil) // Returns false
//	fp.SomeStr(isStringLengthLessThan10, []string{}) // Returns false
//
//	func isStringLengthLessThan10(str string) bool {
//		return len(str) < 10
//	}
func SomeStr(f func(string) bool, list []string) bool {
	if f == nil {
		return false
	}

	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
