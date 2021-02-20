package fp

// SomeIntPtr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeIntPtr(f func(*int) bool, list []*int) bool {
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

// SomeInt64Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeInt64Ptr(f func(*int64) bool, list []*int64) bool {
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

// SomeInt32Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeInt32Ptr(f func(*int32) bool, list []*int32) bool {
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

// SomeInt16Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeInt16Ptr(f func(*int16) bool, list []*int16) bool {
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

// SomeInt8Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeInt8Ptr(f func(*int8) bool, list []*int8) bool {
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

// SomeUintPtr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeUintPtr(f func(*uint) bool, list []*uint) bool {
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

// SomeUint64Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeUint64Ptr(f func(*uint64) bool, list []*uint64) bool {
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

// SomeUint32Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeUint32Ptr(f func(*uint32) bool, list []*uint32) bool {
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

// SomeUint16Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeUint16Ptr(f func(*uint16) bool, list []*uint16) bool {
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

// SomeUint8Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeUint8Ptr(f func(*uint8) bool, list []*uint8) bool {
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

// SomeStrPtr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeStrPtr(f func(*string) bool, list []*string) bool {
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

// SomeBoolPtr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeBoolPtr(f func(*bool) bool, list []*bool) bool {
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

// SomeFloat32Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeFloat32Ptr(f func(*float32) bool, list []*float32) bool {
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

// SomeFloat64Ptr finds item in the list based on supplied function.
//
// Takes 2 input:
//	1. Function
//	2. List
//
// Returns:
//	bool.
//	True if condition satisfies, else false
func SomeFloat64Ptr(f func(*float64) bool, list []*float64) bool {
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
