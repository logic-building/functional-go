package fp

// FilterIntPtr takes two arguments
//  1. Function: takes 1 argument of type int and returns bool
//  2. slice of type []*int
//
// Returns:
//  new filtered list
func FilterIntPtr(f func(*int) bool, list []*int) []*int {
	if f == nil {
		return []*int{}
	}
	var newList []*int
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterInt64Ptr takes two arguments
//  1. Function: takes 1 argument of type int64 and returns bool
//  2. slice of type []*int64
//
// Returns:
//  new filtered list
func FilterInt64Ptr(f func(*int64) bool, list []*int64) []*int64 {
	if f == nil {
		return []*int64{}
	}
	var newList []*int64
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterInt32Ptr takes two arguments
//  1. Function: takes 1 argument of type int32 and returns bool
//  2. slice of type []*int32
//
// Returns:
//  new filtered list
func FilterInt32Ptr(f func(*int32) bool, list []*int32) []*int32 {
	if f == nil {
		return []*int32{}
	}
	var newList []*int32
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterInt16Ptr takes two arguments
//  1. Function: takes 1 argument of type int16 and returns bool
//  2. slice of type []*int16
//
// Returns:
//  new filtered list
func FilterInt16Ptr(f func(*int16) bool, list []*int16) []*int16 {
	if f == nil {
		return []*int16{}
	}
	var newList []*int16
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterInt8Ptr takes two arguments
//  1. Function: takes 1 argument of type int8 and returns bool
//  2. slice of type []*int8
//
// Returns:
//  new filtered list
func FilterInt8Ptr(f func(*int8) bool, list []*int8) []*int8 {
	if f == nil {
		return []*int8{}
	}
	var newList []*int8
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUintPtr takes two arguments
//  1. Function: takes 1 argument of type uint and returns bool
//  2. slice of type []*uint
//
// Returns:
//  new filtered list
func FilterUintPtr(f func(*uint) bool, list []*uint) []*uint {
	if f == nil {
		return []*uint{}
	}
	var newList []*uint
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUint64Ptr takes two arguments
//  1. Function: takes 1 argument of type uint64 and returns bool
//  2. slice of type []*uint64
//
// Returns:
//  new filtered list
func FilterUint64Ptr(f func(*uint64) bool, list []*uint64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	var newList []*uint64
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUint32Ptr takes two arguments
//  1. Function: takes 1 argument of type uint32 and returns bool
//  2. slice of type []*uint32
//
// Returns:
//  new filtered list
func FilterUint32Ptr(f func(*uint32) bool, list []*uint32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	var newList []*uint32
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUint16Ptr takes two arguments
//  1. Function: takes 1 argument of type uint16 and returns bool
//  2. slice of type []*uint16
//
// Returns:
//  new filtered list
func FilterUint16Ptr(f func(*uint16) bool, list []*uint16) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	var newList []*uint16
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterUint8Ptr takes two arguments
//  1. Function: takes 1 argument of type uint8 and returns bool
//  2. slice of type []*uint8
//
// Returns:
//  new filtered list
func FilterUint8Ptr(f func(*uint8) bool, list []*uint8) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	var newList []*uint8
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterStrPtr takes two arguments
//  1. Function: takes 1 argument of type string and returns bool
//  2. slice of type []*string
//
// Returns:
//  new filtered list
func FilterStrPtr(f func(*string) bool, list []*string) []*string {
	if f == nil {
		return []*string{}
	}
	var newList []*string
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterBoolPtr takes two arguments
//  1. Function: takes 1 argument of type bool and returns bool
//  2. slice of type []*bool
//
// Returns:
//  new filtered list
func FilterBoolPtr(f func(*bool) bool, list []*bool) []*bool {
	if f == nil {
		return []*bool{}
	}
	var newList []*bool
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterFloat32Ptr takes two arguments
//  1. Function: takes 1 argument of type float32 and returns bool
//  2. slice of type []*float32
//
// Returns:
//  new filtered list
func FilterFloat32Ptr(f func(*float32) bool, list []*float32) []*float32 {
	if f == nil {
		return []*float32{}
	}
	var newList []*float32
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// FilterFloat64Ptr takes two arguments
//  1. Function: takes 1 argument of type float64 and returns bool
//  2. slice of type []*float64
//
// Returns:
//  new filtered list
func FilterFloat64Ptr(f func(*float64) bool, list []*float64) []*float64 {
	if f == nil {
		return []*float64{}
	}
	var newList []*float64
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
