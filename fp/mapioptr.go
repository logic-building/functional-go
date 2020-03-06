package fp

// MapIntInt64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntInt64Ptr(f func(*int) *int64, list []*int) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntInt32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntInt32Ptr(f func(*int) *int32, list []*int) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntInt16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntInt16Ptr(f func(*int) *int16, list []*int) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntInt8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntInt8Ptr(f func(*int) *int8, list []*int) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntUint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntUintPtr(f func(*int) *uint, list []*int) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntUint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntUint64Ptr(f func(*int) *uint64, list []*int) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntUint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntUint32Ptr(f func(*int) *uint32, list []*int) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntUint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntUint16Ptr(f func(*int) *uint16, list []*int) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntUint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntUint8Ptr(f func(*int) *uint8, list []*int) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntStrPtr(f func(*int) *string, list []*int) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntBool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntBoolPtr(f func(*int) *bool, list []*int) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntFloat32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntFloat32Ptr(f func(*int) *float32, list []*int) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntFloat64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntFloat64Ptr(f func(*int) *float64, list []*int) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64IntPtr(f func(*int64) *int, list []*int64) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64Int32Ptr(f func(*int64) *int32, list []*int64) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64Int16Ptr(f func(*int64) *int16, list []*int64) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64Int8Ptr(f func(*int64) *int8, list []*int64) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64UintPtr(f func(*int64) *uint, list []*int64) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64Uint64Ptr(f func(*int64) *uint64, list []*int64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64Uint32Ptr(f func(*int64) *uint32, list []*int64) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64Uint16Ptr(f func(*int64) *uint16, list []*int64) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64Uint8Ptr(f func(*int64) *uint8, list []*int64) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64StrPtr(f func(*int64) *string, list []*int64) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64BoolPtr(f func(*int64) *bool, list []*int64) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64Float32Ptr(f func(*int64) *float32, list []*int64) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt64Float64Ptr(f func(*int64) *float64, list []*int64) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32IntPtr(f func(*int32) *int, list []*int32) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32Int64Ptr(f func(*int32) *int64, list []*int32) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32Int16Ptr(f func(*int32) *int16, list []*int32) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32Int8Ptr(f func(*int32) *int8, list []*int32) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32UintPtr(f func(*int32) *uint, list []*int32) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32Uint64Ptr(f func(*int32) *uint64, list []*int32) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32Uint32Ptr(f func(*int32) *uint32, list []*int32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32Uint16Ptr(f func(*int32) *uint16, list []*int32) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32Uint8Ptr(f func(*int32) *uint8, list []*int32) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32StrPtr(f func(*int32) *string, list []*int32) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32BoolPtr(f func(*int32) *bool, list []*int32) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32Float32Ptr(f func(*int32) *float32, list []*int32) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt32Float64Ptr(f func(*int32) *float64, list []*int32) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16IntPtr(f func(*int16) *int, list []*int16) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16Int64Ptr(f func(*int16) *int64, list []*int16) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16Int32Ptr(f func(*int16) *int32, list []*int16) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16Int8Ptr(f func(*int16) *int8, list []*int16) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16UintPtr(f func(*int16) *uint, list []*int16) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16Uint64Ptr(f func(*int16) *uint64, list []*int16) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16Uint32Ptr(f func(*int16) *uint32, list []*int16) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16Uint16Ptr(f func(*int16) *uint16, list []*int16) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16Uint8Ptr(f func(*int16) *uint8, list []*int16) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16StrPtr(f func(*int16) *string, list []*int16) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16BoolPtr(f func(*int16) *bool, list []*int16) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16Float32Ptr(f func(*int16) *float32, list []*int16) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt16Float64Ptr(f func(*int16) *float64, list []*int16) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8IntPtr(f func(*int8) *int, list []*int8) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8Int64Ptr(f func(*int8) *int64, list []*int8) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8Int32Ptr(f func(*int8) *int32, list []*int8) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8Int16Ptr(f func(*int8) *int16, list []*int8) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8UintPtr(f func(*int8) *uint, list []*int8) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8Uint64Ptr(f func(*int8) *uint64, list []*int8) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8Uint32Ptr(f func(*int8) *uint32, list []*int8) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8Uint16Ptr(f func(*int8) *uint16, list []*int8) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8Uint8Ptr(f func(*int8) *uint8, list []*int8) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8StrPtr(f func(*int8) *string, list []*int8) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8BoolPtr(f func(*int8) *bool, list []*int8) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8Float32Ptr(f func(*int8) *float32, list []*int8) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInt8Float64Ptr(f func(*int8) *float64, list []*int8) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintIntPtr(f func(*uint) *int, list []*uint) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintInt64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintInt64Ptr(f func(*uint) *int64, list []*uint) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintInt32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintInt32Ptr(f func(*uint) *int32, list []*uint) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintInt16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintInt16Ptr(f func(*uint) *int16, list []*uint) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintInt8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintInt8Ptr(f func(*uint) *int8, list []*uint) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintUint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintUint64Ptr(f func(*uint) *uint64, list []*uint) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintUint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintUint32Ptr(f func(*uint) *uint32, list []*uint) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintUint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintUint16Ptr(f func(*uint) *uint16, list []*uint) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintUint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintUint8Ptr(f func(*uint) *uint8, list []*uint) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintStrPtr(f func(*uint) *string, list []*uint) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintBool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintBoolPtr(f func(*uint) *bool, list []*uint) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintFloat32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintFloat32Ptr(f func(*uint) *float32, list []*uint) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUintFloat64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUintFloat64Ptr(f func(*uint) *float64, list []*uint) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64IntPtr(f func(*uint64) *int, list []*uint64) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64Int64Ptr(f func(*uint64) *int64, list []*uint64) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64Int32Ptr(f func(*uint64) *int32, list []*uint64) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64Int16Ptr(f func(*uint64) *int16, list []*uint64) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64Int8Ptr(f func(*uint64) *int8, list []*uint64) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64UintPtr(f func(*uint64) *uint, list []*uint64) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64Uint32Ptr(f func(*uint64) *uint32, list []*uint64) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64Uint16Ptr(f func(*uint64) *uint16, list []*uint64) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64Uint8Ptr(f func(*uint64) *uint8, list []*uint64) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64StrPtr(f func(*uint64) *string, list []*uint64) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64BoolPtr(f func(*uint64) *bool, list []*uint64) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64Float32Ptr(f func(*uint64) *float32, list []*uint64) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint64Float64Ptr(f func(*uint64) *float64, list []*uint64) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32IntPtr(f func(*uint32) *int, list []*uint32) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32Int64Ptr(f func(*uint32) *int64, list []*uint32) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32Int32Ptr(f func(*uint32) *int32, list []*uint32) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32Int16Ptr(f func(*uint32) *int16, list []*uint32) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32Int8Ptr(f func(*uint32) *int8, list []*uint32) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32UintPtr(f func(*uint32) *uint, list []*uint32) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32Uint64Ptr(f func(*uint32) *uint64, list []*uint32) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32Uint16Ptr(f func(*uint32) *uint16, list []*uint32) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32Uint8Ptr(f func(*uint32) *uint8, list []*uint32) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32StrPtr(f func(*uint32) *string, list []*uint32) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32BoolPtr(f func(*uint32) *bool, list []*uint32) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32Float32Ptr(f func(*uint32) *float32, list []*uint32) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint32Float64Ptr(f func(*uint32) *float64, list []*uint32) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16IntPtr(f func(*uint16) *int, list []*uint16) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16Int64Ptr(f func(*uint16) *int64, list []*uint16) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16Int32Ptr(f func(*uint16) *int32, list []*uint16) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16Int16Ptr(f func(*uint16) *int16, list []*uint16) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16Int8Ptr(f func(*uint16) *int8, list []*uint16) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16UintPtr(f func(*uint16) *uint, list []*uint16) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16Uint64Ptr(f func(*uint16) *uint64, list []*uint16) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16Uint32Ptr(f func(*uint16) *uint32, list []*uint16) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16Uint8Ptr(f func(*uint16) *uint8, list []*uint16) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16StrPtr(f func(*uint16) *string, list []*uint16) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16BoolPtr(f func(*uint16) *bool, list []*uint16) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16Float32Ptr(f func(*uint16) *float32, list []*uint16) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint16Float64Ptr(f func(*uint16) *float64, list []*uint16) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8IntPtr(f func(*uint8) *int, list []*uint8) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8Int64Ptr(f func(*uint8) *int64, list []*uint8) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8Int32Ptr(f func(*uint8) *int32, list []*uint8) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8Int16Ptr(f func(*uint8) *int16, list []*uint8) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8Int8Ptr(f func(*uint8) *int8, list []*uint8) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8UintPtr(f func(*uint8) *uint, list []*uint8) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8Uint64Ptr(f func(*uint8) *uint64, list []*uint8) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8Uint32Ptr(f func(*uint8) *uint32, list []*uint8) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8Uint16Ptr(f func(*uint8) *uint16, list []*uint8) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8StrPtr(f func(*uint8) *string, list []*uint8) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8BoolPtr(f func(*uint8) *bool, list []*uint8) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8Float32Ptr(f func(*uint8) *float32, list []*uint8) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapUint8Float64Ptr(f func(*uint8) *float64, list []*uint8) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrIntPtr(f func(*string) *int, list []*string) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrInt64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrInt64Ptr(f func(*string) *int64, list []*string) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrInt32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrInt32Ptr(f func(*string) *int32, list []*string) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrInt16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrInt16Ptr(f func(*string) *int16, list []*string) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrInt8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrInt8Ptr(f func(*string) *int8, list []*string) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrUint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrUintPtr(f func(*string) *uint, list []*string) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrUint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrUint64Ptr(f func(*string) *uint64, list []*string) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrUint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrUint32Ptr(f func(*string) *uint32, list []*string) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrUint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrUint16Ptr(f func(*string) *uint16, list []*string) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrUint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrUint8Ptr(f func(*string) *uint8, list []*string) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrBool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrBoolPtr(f func(*string) *bool, list []*string) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrFloat32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrFloat32Ptr(f func(*string) *float32, list []*string) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrFloat64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrFloat64Ptr(f func(*string) *float64, list []*string) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolIntPtr(f func(*bool) *int, list []*bool) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolInt64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolInt64Ptr(f func(*bool) *int64, list []*bool) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolInt32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolInt32Ptr(f func(*bool) *int32, list []*bool) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolInt16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolInt16Ptr(f func(*bool) *int16, list []*bool) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolInt8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolInt8Ptr(f func(*bool) *int8, list []*bool) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolUint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolUintPtr(f func(*bool) *uint, list []*bool) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolUint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolUint64Ptr(f func(*bool) *uint64, list []*bool) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolUint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolUint32Ptr(f func(*bool) *uint32, list []*bool) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolUint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolUint16Ptr(f func(*bool) *uint16, list []*bool) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolUint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolUint8Ptr(f func(*bool) *uint8, list []*bool) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolStrPtr(f func(*bool) *string, list []*bool) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolFloat32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolFloat32Ptr(f func(*bool) *float32, list []*bool) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBoolFloat64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapBoolFloat64Ptr(f func(*bool) *float64, list []*bool) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32IntPtr(f func(*float32) *int, list []*float32) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32Int64Ptr(f func(*float32) *int64, list []*float32) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32Int32Ptr(f func(*float32) *int32, list []*float32) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32Int16Ptr(f func(*float32) *int16, list []*float32) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32Int8Ptr(f func(*float32) *int8, list []*float32) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32UintPtr(f func(*float32) *uint, list []*float32) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32Uint64Ptr(f func(*float32) *uint64, list []*float32) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32Uint32Ptr(f func(*float32) *uint32, list []*float32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32Uint16Ptr(f func(*float32) *uint16, list []*float32) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32Uint8Ptr(f func(*float32) *uint8, list []*float32) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32StrPtr(f func(*float32) *string, list []*float32) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32BoolPtr(f func(*float32) *bool, list []*float32) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat32Float64Ptr(f func(*float32) *float64, list []*float32) []*float64 {
	if f == nil {
		return []*float64{}
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64IntPtr(f func(*float64) *int, list []*float64) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64Int64Ptr(f func(*float64) *int64, list []*float64) []*int64 {
	if f == nil {
		return []*int64{}
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64Int32Ptr(f func(*float64) *int32, list []*float64) []*int32 {
	if f == nil {
		return []*int32{}
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64Int16Ptr(f func(*float64) *int16, list []*float64) []*int16 {
	if f == nil {
		return []*int16{}
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64Int8Ptr(f func(*float64) *int8, list []*float64) []*int8 {
	if f == nil {
		return []*int8{}
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64UintPtr(f func(*float64) *uint, list []*float64) []*uint {
	if f == nil {
		return []*uint{}
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64Uint64Ptr(f func(*float64) *uint64, list []*float64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64Uint32Ptr(f func(*float64) *uint32, list []*float64) []*uint32 {
	if f == nil {
		return []*uint32{}
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64Uint16Ptr(f func(*float64) *uint16, list []*float64) []*uint16 {
	if f == nil {
		return []*uint16{}
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64Uint8Ptr(f func(*float64) *uint8, list []*float64) []*uint8 {
	if f == nil {
		return []*uint8{}
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64StrPtr(f func(*float64) *string, list []*float64) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64BoolPtr(f func(*float64) *bool, list []*float64) []*bool {
	if f == nil {
		return []*bool{}
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapFloat64Float32Ptr(f func(*float64) *float32, list []*float64) []*float32 {
	if f == nil {
		return []*float32{}
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
