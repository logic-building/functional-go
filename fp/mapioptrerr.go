package fp

// MapIntInt64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntInt64PtrErr(f func(*int) (*int64, error), list []*int) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntInt32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntInt32PtrErr(f func(*int) (*int32, error), list []*int) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntInt16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntInt16PtrErr(f func(*int) (*int16, error), list []*int) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntInt8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntInt8PtrErr(f func(*int) (*int8, error), list []*int) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntUintPtrErr(f func(*int) (*uint, error), list []*int) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntUint64PtrErr(f func(*int) (*uint64, error), list []*int) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntUint32PtrErr(f func(*int) (*uint32, error), list []*int) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntUint16PtrErr(f func(*int) (*uint16, error), list []*int) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntUint8PtrErr(f func(*int) (*uint8, error), list []*int) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntStrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntStrPtrErr(f func(*int) (*string, error), list []*int) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntBoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntBoolPtrErr(f func(*int) (*bool, error), list []*int) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntFloat32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntFloat32PtrErr(f func(*int) (*float32, error), list []*int) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntFloat64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapIntFloat64PtrErr(f func(*int) (*float64, error), list []*int) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64IntPtrErr(f func(*int64) (*int, error), list []*int64) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Int32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64Int32PtrErr(f func(*int64) (*int32, error), list []*int64) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Int16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64Int16PtrErr(f func(*int64) (*int16, error), list []*int64) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Int8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64Int8PtrErr(f func(*int64) (*int8, error), list []*int64) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64UintPtrErr(f func(*int64) (*uint, error), list []*int64) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Uint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64Uint64PtrErr(f func(*int64) (*uint64, error), list []*int64) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Uint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64Uint32PtrErr(f func(*int64) (*uint32, error), list []*int64) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Uint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64Uint16PtrErr(f func(*int64) (*uint16, error), list []*int64) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Uint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64Uint8PtrErr(f func(*int64) (*uint8, error), list []*int64) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64StrPtrErr(f func(*int64) (*string, error), list []*int64) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64BoolPtrErr(f func(*int64) (*bool, error), list []*int64) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Float32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64Float32PtrErr(f func(*int64) (*float32, error), list []*int64) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Float64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt64Float64PtrErr(f func(*int64) (*float64, error), list []*int64) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32IntPtrErr(f func(*int32) (*int, error), list []*int32) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Int64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32Int64PtrErr(f func(*int32) (*int64, error), list []*int32) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Int16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32Int16PtrErr(f func(*int32) (*int16, error), list []*int32) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Int8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32Int8PtrErr(f func(*int32) (*int8, error), list []*int32) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32UintPtrErr(f func(*int32) (*uint, error), list []*int32) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Uint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32Uint64PtrErr(f func(*int32) (*uint64, error), list []*int32) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Uint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32Uint32PtrErr(f func(*int32) (*uint32, error), list []*int32) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Uint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32Uint16PtrErr(f func(*int32) (*uint16, error), list []*int32) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Uint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32Uint8PtrErr(f func(*int32) (*uint8, error), list []*int32) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32StrPtrErr(f func(*int32) (*string, error), list []*int32) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32BoolPtrErr(f func(*int32) (*bool, error), list []*int32) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Float32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32Float32PtrErr(f func(*int32) (*float32, error), list []*int32) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Float64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt32Float64PtrErr(f func(*int32) (*float64, error), list []*int32) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16IntPtrErr(f func(*int16) (*int, error), list []*int16) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Int64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16Int64PtrErr(f func(*int16) (*int64, error), list []*int16) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Int32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16Int32PtrErr(f func(*int16) (*int32, error), list []*int16) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Int8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16Int8PtrErr(f func(*int16) (*int8, error), list []*int16) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16UintPtrErr(f func(*int16) (*uint, error), list []*int16) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Uint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16Uint64PtrErr(f func(*int16) (*uint64, error), list []*int16) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Uint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16Uint32PtrErr(f func(*int16) (*uint32, error), list []*int16) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Uint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16Uint16PtrErr(f func(*int16) (*uint16, error), list []*int16) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Uint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16Uint8PtrErr(f func(*int16) (*uint8, error), list []*int16) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16StrPtrErr(f func(*int16) (*string, error), list []*int16) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16BoolPtrErr(f func(*int16) (*bool, error), list []*int16) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Float32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16Float32PtrErr(f func(*int16) (*float32, error), list []*int16) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Float64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt16Float64PtrErr(f func(*int16) (*float64, error), list []*int16) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8IntPtrErr(f func(*int8) (*int, error), list []*int8) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Int64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8Int64PtrErr(f func(*int8) (*int64, error), list []*int8) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Int32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8Int32PtrErr(f func(*int8) (*int32, error), list []*int8) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Int16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8Int16PtrErr(f func(*int8) (*int16, error), list []*int8) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8UintPtrErr(f func(*int8) (*uint, error), list []*int8) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Uint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8Uint64PtrErr(f func(*int8) (*uint64, error), list []*int8) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Uint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8Uint32PtrErr(f func(*int8) (*uint32, error), list []*int8) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Uint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8Uint16PtrErr(f func(*int8) (*uint16, error), list []*int8) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Uint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8Uint8PtrErr(f func(*int8) (*uint8, error), list []*int8) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8StrPtrErr(f func(*int8) (*string, error), list []*int8) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8BoolPtrErr(f func(*int8) (*bool, error), list []*int8) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Float32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8Float32PtrErr(f func(*int8) (*float32, error), list []*int8) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Float64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapInt8Float64PtrErr(f func(*int8) (*float64, error), list []*int8) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintIntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintIntPtrErr(f func(*uint) (*int, error), list []*uint) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintInt64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintInt64PtrErr(f func(*uint) (*int64, error), list []*uint) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintInt32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintInt32PtrErr(f func(*uint) (*int32, error), list []*uint) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintInt16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintInt16PtrErr(f func(*uint) (*int16, error), list []*uint) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintInt8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintInt8PtrErr(f func(*uint) (*int8, error), list []*uint) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintUint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintUint64PtrErr(f func(*uint) (*uint64, error), list []*uint) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintUint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintUint32PtrErr(f func(*uint) (*uint32, error), list []*uint) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintUint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintUint16PtrErr(f func(*uint) (*uint16, error), list []*uint) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintUint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintUint8PtrErr(f func(*uint) (*uint8, error), list []*uint) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintStrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintStrPtrErr(f func(*uint) (*string, error), list []*uint) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintBoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintBoolPtrErr(f func(*uint) (*bool, error), list []*uint) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintFloat32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintFloat32PtrErr(f func(*uint) (*float32, error), list []*uint) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintFloat64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUintFloat64PtrErr(f func(*uint) (*float64, error), list []*uint) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64IntPtrErr(f func(*uint64) (*int, error), list []*uint64) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Int64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64Int64PtrErr(f func(*uint64) (*int64, error), list []*uint64) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Int32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64Int32PtrErr(f func(*uint64) (*int32, error), list []*uint64) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Int16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64Int16PtrErr(f func(*uint64) (*int16, error), list []*uint64) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Int8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64Int8PtrErr(f func(*uint64) (*int8, error), list []*uint64) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64UintPtrErr(f func(*uint64) (*uint, error), list []*uint64) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Uint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64Uint32PtrErr(f func(*uint64) (*uint32, error), list []*uint64) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Uint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64Uint16PtrErr(f func(*uint64) (*uint16, error), list []*uint64) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Uint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64Uint8PtrErr(f func(*uint64) (*uint8, error), list []*uint64) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64StrPtrErr(f func(*uint64) (*string, error), list []*uint64) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64BoolPtrErr(f func(*uint64) (*bool, error), list []*uint64) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Float32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64Float32PtrErr(f func(*uint64) (*float32, error), list []*uint64) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Float64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint64Float64PtrErr(f func(*uint64) (*float64, error), list []*uint64) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32IntPtrErr(f func(*uint32) (*int, error), list []*uint32) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Int64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32Int64PtrErr(f func(*uint32) (*int64, error), list []*uint32) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Int32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32Int32PtrErr(f func(*uint32) (*int32, error), list []*uint32) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Int16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32Int16PtrErr(f func(*uint32) (*int16, error), list []*uint32) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Int8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32Int8PtrErr(f func(*uint32) (*int8, error), list []*uint32) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32UintPtrErr(f func(*uint32) (*uint, error), list []*uint32) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Uint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32Uint64PtrErr(f func(*uint32) (*uint64, error), list []*uint32) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Uint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32Uint16PtrErr(f func(*uint32) (*uint16, error), list []*uint32) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Uint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32Uint8PtrErr(f func(*uint32) (*uint8, error), list []*uint32) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32StrPtrErr(f func(*uint32) (*string, error), list []*uint32) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32BoolPtrErr(f func(*uint32) (*bool, error), list []*uint32) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Float32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32Float32PtrErr(f func(*uint32) (*float32, error), list []*uint32) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Float64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint32Float64PtrErr(f func(*uint32) (*float64, error), list []*uint32) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16IntPtrErr(f func(*uint16) (*int, error), list []*uint16) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Int64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16Int64PtrErr(f func(*uint16) (*int64, error), list []*uint16) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Int32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16Int32PtrErr(f func(*uint16) (*int32, error), list []*uint16) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Int16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16Int16PtrErr(f func(*uint16) (*int16, error), list []*uint16) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Int8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16Int8PtrErr(f func(*uint16) (*int8, error), list []*uint16) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16UintPtrErr(f func(*uint16) (*uint, error), list []*uint16) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Uint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16Uint64PtrErr(f func(*uint16) (*uint64, error), list []*uint16) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Uint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16Uint32PtrErr(f func(*uint16) (*uint32, error), list []*uint16) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Uint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16Uint8PtrErr(f func(*uint16) (*uint8, error), list []*uint16) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16StrPtrErr(f func(*uint16) (*string, error), list []*uint16) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16BoolPtrErr(f func(*uint16) (*bool, error), list []*uint16) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Float32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16Float32PtrErr(f func(*uint16) (*float32, error), list []*uint16) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Float64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint16Float64PtrErr(f func(*uint16) (*float64, error), list []*uint16) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8IntPtrErr(f func(*uint8) (*int, error), list []*uint8) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Int64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8Int64PtrErr(f func(*uint8) (*int64, error), list []*uint8) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Int32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8Int32PtrErr(f func(*uint8) (*int32, error), list []*uint8) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Int16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8Int16PtrErr(f func(*uint8) (*int16, error), list []*uint8) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Int8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8Int8PtrErr(f func(*uint8) (*int8, error), list []*uint8) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8UintPtrErr(f func(*uint8) (*uint, error), list []*uint8) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Uint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8Uint64PtrErr(f func(*uint8) (*uint64, error), list []*uint8) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Uint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8Uint32PtrErr(f func(*uint8) (*uint32, error), list []*uint8) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Uint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8Uint16PtrErr(f func(*uint8) (*uint16, error), list []*uint8) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8StrPtrErr(f func(*uint8) (*string, error), list []*uint8) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8BoolPtrErr(f func(*uint8) (*bool, error), list []*uint8) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Float32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8Float32PtrErr(f func(*uint8) (*float32, error), list []*uint8) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Float64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapUint8Float64PtrErr(f func(*uint8) (*float64, error), list []*uint8) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrIntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrIntPtrErr(f func(*string) (*int, error), list []*string) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrInt64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrInt64PtrErr(f func(*string) (*int64, error), list []*string) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrInt32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrInt32PtrErr(f func(*string) (*int32, error), list []*string) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrInt16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrInt16PtrErr(f func(*string) (*int16, error), list []*string) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrInt8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrInt8PtrErr(f func(*string) (*int8, error), list []*string) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrUintPtrErr(f func(*string) (*uint, error), list []*string) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrUint64PtrErr(f func(*string) (*uint64, error), list []*string) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrUint32PtrErr(f func(*string) (*uint32, error), list []*string) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrUint16PtrErr(f func(*string) (*uint16, error), list []*string) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrUint8PtrErr(f func(*string) (*uint8, error), list []*string) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrBoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrBoolPtrErr(f func(*string) (*bool, error), list []*string) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrFloat32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrFloat32PtrErr(f func(*string) (*float32, error), list []*string) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrFloat64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapStrFloat64PtrErr(f func(*string) (*float64, error), list []*string) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolIntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolIntPtrErr(f func(*bool) (*int, error), list []*bool) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolInt64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolInt64PtrErr(f func(*bool) (*int64, error), list []*bool) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolInt32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolInt32PtrErr(f func(*bool) (*int32, error), list []*bool) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolInt16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolInt16PtrErr(f func(*bool) (*int16, error), list []*bool) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolInt8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolInt8PtrErr(f func(*bool) (*int8, error), list []*bool) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolUintPtrErr(f func(*bool) (*uint, error), list []*bool) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolUint64PtrErr(f func(*bool) (*uint64, error), list []*bool) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolUint32PtrErr(f func(*bool) (*uint32, error), list []*bool) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolUint16PtrErr(f func(*bool) (*uint16, error), list []*bool) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolUint8PtrErr(f func(*bool) (*uint8, error), list []*bool) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolStrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolStrPtrErr(f func(*bool) (*string, error), list []*bool) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolFloat32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolFloat32PtrErr(f func(*bool) (*float32, error), list []*bool) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolFloat64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapBoolFloat64PtrErr(f func(*bool) (*float64, error), list []*bool) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32IntPtrErr(f func(*float32) (*int, error), list []*float32) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Int64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32Int64PtrErr(f func(*float32) (*int64, error), list []*float32) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Int32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32Int32PtrErr(f func(*float32) (*int32, error), list []*float32) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Int16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32Int16PtrErr(f func(*float32) (*int16, error), list []*float32) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Int8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32Int8PtrErr(f func(*float32) (*int8, error), list []*float32) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32UintPtrErr(f func(*float32) (*uint, error), list []*float32) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Uint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32Uint64PtrErr(f func(*float32) (*uint64, error), list []*float32) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Uint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32Uint32PtrErr(f func(*float32) (*uint32, error), list []*float32) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Uint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32Uint16PtrErr(f func(*float32) (*uint16, error), list []*float32) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Uint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32Uint8PtrErr(f func(*float32) (*uint8, error), list []*float32) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32StrPtrErr(f func(*float32) (*string, error), list []*float32) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32BoolPtrErr(f func(*float32) (*bool, error), list []*float32) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Float64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat32Float64PtrErr(f func(*float32) (*float64, error), list []*float32) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}
	newList := make([]*float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64IntPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64IntPtrErr(f func(*float64) (*int, error), list []*float64) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Int64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64Int64PtrErr(f func(*float64) (*int64, error), list []*float64) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}
	newList := make([]*int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Int32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64Int32PtrErr(f func(*float64) (*int32, error), list []*float64) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}
	newList := make([]*int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Int16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64Int16PtrErr(f func(*float64) (*int16, error), list []*float64) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}
	newList := make([]*int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Int8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64Int8PtrErr(f func(*float64) (*int8, error), list []*float64) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}
	newList := make([]*int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64UintPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64UintPtrErr(f func(*float64) (*uint, error), list []*float64) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}
	newList := make([]*uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Uint64PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64Uint64PtrErr(f func(*float64) (*uint64, error), list []*float64) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}
	newList := make([]*uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Uint32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64Uint32PtrErr(f func(*float64) (*uint32, error), list []*float64) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}
	newList := make([]*uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Uint16PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64Uint16PtrErr(f func(*float64) (*uint16, error), list []*float64) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}
	newList := make([]*uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Uint8PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64Uint8PtrErr(f func(*float64) (*uint8, error), list []*float64) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}
	newList := make([]*uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64StrPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64StrPtrErr(f func(*float64) (*string, error), list []*float64) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64BoolPtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64BoolPtrErr(f func(*float64) (*bool, error), list []*float64) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}
	newList := make([]*bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Float32PtrErr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list, error
func MapFloat64Float32PtrErr(f func(*float64) (*float32, error), list []*float64) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}
	newList := make([]*float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}
