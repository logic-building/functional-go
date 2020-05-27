package fp

// MapIntPtr takes 2 arguments.
//  1. A function input argument: int and return type (int)
//  2. A list of type []*int
//
// Returns:
// 	([]*int)
func MapIntPtr(f func(*int) *int, list []*int) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Ptr takes 2 arguments.
//  1. A function input argument: int64 and return type (int64)
//  2. A list of type []*int64
//
// Returns:
// 	([]*int64)
func MapInt64Ptr(f func(*int64) *int64, list []*int64) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Ptr takes 2 arguments.
//  1. A function input argument: int32 and return type (int32)
//  2. A list of type []*int32
//
// Returns:
// 	([]*int32)
func MapInt32Ptr(f func(*int32) *int32, list []*int32) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Ptr takes 2 arguments.
//  1. A function input argument: int16 and return type (int16)
//  2. A list of type []*int16
//
// Returns:
// 	([]*int16)
func MapInt16Ptr(f func(*int16) *int16, list []*int16) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Ptr takes 2 arguments.
//  1. A function input argument: int8 and return type (int8)
//  2. A list of type []*int8
//
// Returns:
// 	([]*int8)
func MapInt8Ptr(f func(*int8) *int8, list []*int8) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintPtr takes 2 arguments.
//  1. A function input argument: uint and return type (uint)
//  2. A list of type []*uint
//
// Returns:
// 	([]*uint)
func MapUintPtr(f func(*uint) *uint, list []*uint) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Ptr takes 2 arguments.
//  1. A function input argument: uint64 and return type (uint64)
//  2. A list of type []*uint64
//
// Returns:
// 	([]*uint64)
func MapUint64Ptr(f func(*uint64) *uint64, list []*uint64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Ptr takes 2 arguments.
//  1. A function input argument: uint32 and return type (uint32)
//  2. A list of type []*uint32
//
// Returns:
// 	([]*uint32)
func MapUint32Ptr(f func(*uint32) *uint32, list []*uint32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Ptr takes 2 arguments.
//  1. A function input argument: uint16 and return type (uint16)
//  2. A list of type []*uint16
//
// Returns:
// 	([]*uint16)
func MapUint16Ptr(f func(*uint16) *uint16, list []*uint16) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Ptr takes 2 arguments.
//  1. A function input argument: uint8 and return type (uint8)
//  2. A list of type []*uint8
//
// Returns:
// 	([]*uint8)
func MapUint8Ptr(f func(*uint8) *uint8, list []*uint8) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrPtr takes 2 arguments.
//  1. A function input argument: string and return type (string)
//  2. A list of type []*string
//
// Returns:
// 	([]*string)
func MapStrPtr(f func(*string) *string, list []*string) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolPtr takes 2 arguments.
//  1. A function input argument: bool and return type (bool)
//  2. A list of type []*bool
//
// Returns:
// 	([]*bool)
func MapBoolPtr(f func(*bool) *bool, list []*bool) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Ptr takes 2 arguments.
//  1. A function input argument: float32 and return type (float32)
//  2. A list of type []*float32
//
// Returns:
// 	([]*float32)
func MapFloat32Ptr(f func(*float32) *float32, list []*float32) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Ptr takes 2 arguments.
//  1. A function input argument: float64 and return type (float64)
//  2. A list of type []*float64
//
// Returns:
// 	([]*float64)
func MapFloat64Ptr(f func(*float64) *float64, list []*float64) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
