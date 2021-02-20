package fp

// MapIntErr takes 2 arguments:
//  1. A function input argument: int and return types (int, error)
//  2. A list of type []int
//
// Returns:
// 	([]int, error)
func MapIntErr(f func(int) (int, error), list []int) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	newList := make([]int, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapInt64Err takes 2 arguments:
//  1. A function input argument: int64 and return types (int64, error)
//  2. A list of type []int64
//
// Returns:
// 	([]int64, error)
func MapInt64Err(f func(int64) (int64, error), list []int64) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapInt32Err takes 2 arguments:
//  1. A function input argument: int32 and return types (int32, error)
//  2. A list of type []int32
//
// Returns:
// 	([]int32, error)
func MapInt32Err(f func(int32) (int32, error), list []int32) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapInt16Err takes 2 arguments:
//  1. A function input argument: int16 and return types (int16, error)
//  2. A list of type []int16
//
// Returns:
// 	([]int16, error)
func MapInt16Err(f func(int16) (int16, error), list []int16) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapInt8Err takes 2 arguments:
//  1. A function input argument: int8 and return types (int8, error)
//  2. A list of type []int8
//
// Returns:
// 	([]int8, error)
func MapInt8Err(f func(int8) (int8, error), list []int8) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapUintErr takes 2 arguments:
//  1. A function input argument: uint and return types (uint, error)
//  2. A list of type []uint
//
// Returns:
// 	([]uint, error)
func MapUintErr(f func(uint) (uint, error), list []uint) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapUint64Err takes 2 arguments:
//  1. A function input argument: uint64 and return types (uint64, error)
//  2. A list of type []uint64
//
// Returns:
// 	([]uint64, error)
func MapUint64Err(f func(uint64) (uint64, error), list []uint64) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapUint32Err takes 2 arguments:
//  1. A function input argument: uint32 and return types (uint32, error)
//  2. A list of type []uint32
//
// Returns:
// 	([]uint32, error)
func MapUint32Err(f func(uint32) (uint32, error), list []uint32) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapUint16Err takes 2 arguments:
//  1. A function input argument: uint16 and return types (uint16, error)
//  2. A list of type []uint16
//
// Returns:
// 	([]uint16, error)
func MapUint16Err(f func(uint16) (uint16, error), list []uint16) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapUint8Err takes 2 arguments:
//  1. A function input argument: uint8 and return types (uint8, error)
//  2. A list of type []uint8
//
// Returns:
// 	([]uint8, error)
func MapUint8Err(f func(uint8) (uint8, error), list []uint8) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapStrErr takes 2 arguments:
//  1. A function input argument: string and return types (string, error)
//  2. A list of type []string
//
// Returns:
// 	([]string, error)
func MapStrErr(f func(string) (string, error), list []string) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	newList := make([]string, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapBoolErr takes 2 arguments:
//  1. A function input argument: bool and return types (bool, error)
//  2. A list of type []bool
//
// Returns:
// 	([]bool, error)
func MapBoolErr(f func(bool) (bool, error), list []bool) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapFloat32Err takes 2 arguments:
//  1. A function input argument: float32 and return types (float32, error)
//  2. A list of type []float32
//
// Returns:
// 	([]float32, error)
func MapFloat32Err(f func(float32) (float32, error), list []float32) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}

// MapFloat64Err takes 2 arguments:
//  1. A function input argument: float64 and return types (float64, error)
//  2. A list of type []float64
//
// Returns:
// 	([]float64, error)
func MapFloat64Err(f func(float64) (float64, error), list []float64) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		var err error
		newList[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newList, nil
}
