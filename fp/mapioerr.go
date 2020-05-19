package fp

// MapIntInt64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntInt64Err(f func(int) (int64, error), list []int) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntInt32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntInt32Err(f func(int) (int32, error), list []int) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntInt16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntInt16Err(f func(int) (int16, error), list []int) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntInt8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntInt8Err(f func(int) (int8, error), list []int) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntUintErr(f func(int) (uint, error), list []int) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntUint64Err(f func(int) (uint64, error), list []int) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntUint32Err(f func(int) (uint32, error), list []int) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntUint16Err(f func(int) (uint16, error), list []int) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntUint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntUint8Err(f func(int) (uint8, error), list []int) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntStrErr(f func(int) (string, error), list []int) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntBool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntBoolErr(f func(int) (bool, error), list []int) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntFloat32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntFloat32Err(f func(int) (float32, error), list []int) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapIntFloat64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapIntFloat64Err(f func(int) (float64, error), list []int) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64IntErr(f func(int64) (int, error), list []int64) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64Int32Err(f func(int64) (int32, error), list []int64) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64Int16Err(f func(int64) (int16, error), list []int64) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64Int8Err(f func(int64) (int8, error), list []int64) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64UintErr(f func(int64) (uint, error), list []int64) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64Uint64Err(f func(int64) (uint64, error), list []int64) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64Uint32Err(f func(int64) (uint32, error), list []int64) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64Uint16Err(f func(int64) (uint16, error), list []int64) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64Uint8Err(f func(int64) (uint8, error), list []int64) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64StrErr(f func(int64) (string, error), list []int64) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64BoolErr(f func(int64) (bool, error), list []int64) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64Float32Err(f func(int64) (float32, error), list []int64) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt64Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt64Float64Err(f func(int64) (float64, error), list []int64) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32IntErr(f func(int32) (int, error), list []int32) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32Int64Err(f func(int32) (int64, error), list []int32) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32Int16Err(f func(int32) (int16, error), list []int32) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32Int8Err(f func(int32) (int8, error), list []int32) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32UintErr(f func(int32) (uint, error), list []int32) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32Uint64Err(f func(int32) (uint64, error), list []int32) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32Uint32Err(f func(int32) (uint32, error), list []int32) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32Uint16Err(f func(int32) (uint16, error), list []int32) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32Uint8Err(f func(int32) (uint8, error), list []int32) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32StrErr(f func(int32) (string, error), list []int32) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32BoolErr(f func(int32) (bool, error), list []int32) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32Float32Err(f func(int32) (float32, error), list []int32) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt32Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt32Float64Err(f func(int32) (float64, error), list []int32) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16IntErr(f func(int16) (int, error), list []int16) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16Int64Err(f func(int16) (int64, error), list []int16) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16Int32Err(f func(int16) (int32, error), list []int16) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16Int8Err(f func(int16) (int8, error), list []int16) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16UintErr(f func(int16) (uint, error), list []int16) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16Uint64Err(f func(int16) (uint64, error), list []int16) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16Uint32Err(f func(int16) (uint32, error), list []int16) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16Uint16Err(f func(int16) (uint16, error), list []int16) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16Uint8Err(f func(int16) (uint8, error), list []int16) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16StrErr(f func(int16) (string, error), list []int16) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16BoolErr(f func(int16) (bool, error), list []int16) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16Float32Err(f func(int16) (float32, error), list []int16) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt16Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt16Float64Err(f func(int16) (float64, error), list []int16) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8IntErr(f func(int8) (int, error), list []int8) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8Int64Err(f func(int8) (int64, error), list []int8) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8Int32Err(f func(int8) (int32, error), list []int8) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8Int16Err(f func(int8) (int16, error), list []int8) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8UintErr(f func(int8) (uint, error), list []int8) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8Uint64Err(f func(int8) (uint64, error), list []int8) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8Uint32Err(f func(int8) (uint32, error), list []int8) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8Uint16Err(f func(int8) (uint16, error), list []int8) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8Uint8Err(f func(int8) (uint8, error), list []int8) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8StrErr(f func(int8) (string, error), list []int8) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8BoolErr(f func(int8) (bool, error), list []int8) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8Float32Err(f func(int8) (float32, error), list []int8) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapInt8Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapInt8Float64Err(f func(int8) (float64, error), list []int8) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintIntErr(f func(uint) (int, error), list []uint) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintInt64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintInt64Err(f func(uint) (int64, error), list []uint) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintInt32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintInt32Err(f func(uint) (int32, error), list []uint) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintInt16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintInt16Err(f func(uint) (int16, error), list []uint) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintInt8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintInt8Err(f func(uint) (int8, error), list []uint) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintUint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintUint64Err(f func(uint) (uint64, error), list []uint) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintUint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintUint32Err(f func(uint) (uint32, error), list []uint) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintUint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintUint16Err(f func(uint) (uint16, error), list []uint) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintUint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintUint8Err(f func(uint) (uint8, error), list []uint) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintStrErr(f func(uint) (string, error), list []uint) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintBool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintBoolErr(f func(uint) (bool, error), list []uint) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintFloat32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintFloat32Err(f func(uint) (float32, error), list []uint) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUintFloat64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUintFloat64Err(f func(uint) (float64, error), list []uint) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64IntErr(f func(uint64) (int, error), list []uint64) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64Int64Err(f func(uint64) (int64, error), list []uint64) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64Int32Err(f func(uint64) (int32, error), list []uint64) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64Int16Err(f func(uint64) (int16, error), list []uint64) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64Int8Err(f func(uint64) (int8, error), list []uint64) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64UintErr(f func(uint64) (uint, error), list []uint64) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64Uint32Err(f func(uint64) (uint32, error), list []uint64) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64Uint16Err(f func(uint64) (uint16, error), list []uint64) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64Uint8Err(f func(uint64) (uint8, error), list []uint64) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64StrErr(f func(uint64) (string, error), list []uint64) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64BoolErr(f func(uint64) (bool, error), list []uint64) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64Float32Err(f func(uint64) (float32, error), list []uint64) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint64Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint64Float64Err(f func(uint64) (float64, error), list []uint64) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32IntErr(f func(uint32) (int, error), list []uint32) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32Int64Err(f func(uint32) (int64, error), list []uint32) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32Int32Err(f func(uint32) (int32, error), list []uint32) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32Int16Err(f func(uint32) (int16, error), list []uint32) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32Int8Err(f func(uint32) (int8, error), list []uint32) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32UintErr(f func(uint32) (uint, error), list []uint32) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32Uint64Err(f func(uint32) (uint64, error), list []uint32) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32Uint16Err(f func(uint32) (uint16, error), list []uint32) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32Uint8Err(f func(uint32) (uint8, error), list []uint32) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32StrErr(f func(uint32) (string, error), list []uint32) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32BoolErr(f func(uint32) (bool, error), list []uint32) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32Float32Err(f func(uint32) (float32, error), list []uint32) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint32Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint32Float64Err(f func(uint32) (float64, error), list []uint32) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16IntErr(f func(uint16) (int, error), list []uint16) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16Int64Err(f func(uint16) (int64, error), list []uint16) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16Int32Err(f func(uint16) (int32, error), list []uint16) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16Int16Err(f func(uint16) (int16, error), list []uint16) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16Int8Err(f func(uint16) (int8, error), list []uint16) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16UintErr(f func(uint16) (uint, error), list []uint16) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16Uint64Err(f func(uint16) (uint64, error), list []uint16) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16Uint32Err(f func(uint16) (uint32, error), list []uint16) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16Uint8Err(f func(uint16) (uint8, error), list []uint16) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16StrErr(f func(uint16) (string, error), list []uint16) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16BoolErr(f func(uint16) (bool, error), list []uint16) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16Float32Err(f func(uint16) (float32, error), list []uint16) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint16Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint16Float64Err(f func(uint16) (float64, error), list []uint16) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8IntErr(f func(uint8) (int, error), list []uint8) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8Int64Err(f func(uint8) (int64, error), list []uint8) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8Int32Err(f func(uint8) (int32, error), list []uint8) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8Int16Err(f func(uint8) (int16, error), list []uint8) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8Int8Err(f func(uint8) (int8, error), list []uint8) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8UintErr(f func(uint8) (uint, error), list []uint8) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8Uint64Err(f func(uint8) (uint64, error), list []uint8) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8Uint32Err(f func(uint8) (uint32, error), list []uint8) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8Uint16Err(f func(uint8) (uint16, error), list []uint8) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8StrErr(f func(uint8) (string, error), list []uint8) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8BoolErr(f func(uint8) (bool, error), list []uint8) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8Float32Err(f func(uint8) (float32, error), list []uint8) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapUint8Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapUint8Float64Err(f func(uint8) (float64, error), list []uint8) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrIntErr(f func(string) (int, error), list []string) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrInt64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrInt64Err(f func(string) (int64, error), list []string) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrInt32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrInt32Err(f func(string) (int32, error), list []string) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrInt16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrInt16Err(f func(string) (int16, error), list []string) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrInt8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrInt8Err(f func(string) (int8, error), list []string) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrUintErr(f func(string) (uint, error), list []string) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrUint64Err(f func(string) (uint64, error), list []string) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrUint32Err(f func(string) (uint32, error), list []string) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrUint16Err(f func(string) (uint16, error), list []string) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrUint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrUint8Err(f func(string) (uint8, error), list []string) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrBool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrBoolErr(f func(string) (bool, error), list []string) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrFloat32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrFloat32Err(f func(string) (float32, error), list []string) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapStrFloat64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapStrFloat64Err(f func(string) (float64, error), list []string) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolIntErr(f func(bool) (int, error), list []bool) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolInt64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolInt64Err(f func(bool) (int64, error), list []bool) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolInt32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolInt32Err(f func(bool) (int32, error), list []bool) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolInt16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolInt16Err(f func(bool) (int16, error), list []bool) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolInt8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolInt8Err(f func(bool) (int8, error), list []bool) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolUintErr(f func(bool) (uint, error), list []bool) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolUint64Err(f func(bool) (uint64, error), list []bool) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolUint32Err(f func(bool) (uint32, error), list []bool) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolUint16Err(f func(bool) (uint16, error), list []bool) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolUint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolUint8Err(f func(bool) (uint8, error), list []bool) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolStrErr(f func(bool) (string, error), list []bool) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolFloat32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolFloat32Err(f func(bool) (float32, error), list []bool) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapBoolFloat64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapBoolFloat64Err(f func(bool) (float64, error), list []bool) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32IntErr(f func(float32) (int, error), list []float32) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32Int64Err(f func(float32) (int64, error), list []float32) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32Int32Err(f func(float32) (int32, error), list []float32) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32Int16Err(f func(float32) (int16, error), list []float32) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32Int8Err(f func(float32) (int8, error), list []float32) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32UintErr(f func(float32) (uint, error), list []float32) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32Uint64Err(f func(float32) (uint64, error), list []float32) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32Uint32Err(f func(float32) (uint32, error), list []float32) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32Uint16Err(f func(float32) (uint16, error), list []float32) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32Uint8Err(f func(float32) (uint8, error), list []float32) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32StrErr(f func(float32) (string, error), list []float32) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32BoolErr(f func(float32) (bool, error), list []float32) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat32Float64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat32Float64Err(f func(float32) (float64, error), list []float32) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Int takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64IntErr(f func(float64) (int, error), list []float64) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Int64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64Int64Err(f func(float64) (int64, error), list []float64) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Int32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64Int32Err(f func(float64) (int32, error), list []float64) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Int16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64Int16Err(f func(float64) (int16, error), list []float64) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Int8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64Int8Err(f func(float64) (int8, error), list []float64) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Uint takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64UintErr(f func(float64) (uint, error), list []float64) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Uint64 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64Uint64Err(f func(float64) (uint64, error), list []float64) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Uint32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64Uint32Err(f func(float64) (uint32, error), list []float64) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Uint16 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64Uint16Err(f func(float64) (uint16, error), list []float64) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Uint8 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64Uint8Err(f func(float64) (uint8, error), list []float64) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Str takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64StrErr(f func(float64) (string, error), list []float64) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Bool takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64BoolErr(f func(float64) (bool, error), list []float64) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}

// MapFloat64Float32 takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list and error
func MapFloat64Float32Err(f func(float64) (float32, error), list []float64) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		newList[i] = r
	}
	return newList, nil
}
