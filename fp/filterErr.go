package fp

// FilterIntErr takes two arguments
//  1. Function: takes 1 argument of type int and returns (bool, error)
//  2. slice of type []int
//
// Returns:
//  new filtered list and error
func FilterIntErr(f func(int) (bool, error), list []int) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	var newList []int
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterInt64Err takes two arguments
//  1. Function: takes 1 argument of type int64 and returns (bool, error)
//  2. slice of type []int64
//
// Returns:
//  new filtered list and error
func FilterInt64Err(f func(int64) (bool, error), list []int64) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	var newList []int64
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterInt32Err takes two arguments
//  1. Function: takes 1 argument of type int32 and returns (bool, error)
//  2. slice of type []int32
//
// Returns:
//  new filtered list and error
func FilterInt32Err(f func(int32) (bool, error), list []int32) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	var newList []int32
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterInt16Err takes two arguments
//  1. Function: takes 1 argument of type int16 and returns (bool, error)
//  2. slice of type []int16
//
// Returns:
//  new filtered list and error
func FilterInt16Err(f func(int16) (bool, error), list []int16) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	var newList []int16
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterInt8Err takes two arguments
//  1. Function: takes 1 argument of type int8 and returns (bool, error)
//  2. slice of type []int8
//
// Returns:
//  new filtered list and error
func FilterInt8Err(f func(int8) (bool, error), list []int8) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	var newList []int8
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterUintErr takes two arguments
//  1. Function: takes 1 argument of type uint and returns (bool, error)
//  2. slice of type []uint
//
// Returns:
//  new filtered list and error
func FilterUintErr(f func(uint) (bool, error), list []uint) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	var newList []uint
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterUint64Err takes two arguments
//  1. Function: takes 1 argument of type uint64 and returns (bool, error)
//  2. slice of type []uint64
//
// Returns:
//  new filtered list and error
func FilterUint64Err(f func(uint64) (bool, error), list []uint64) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	var newList []uint64
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterUint32Err takes two arguments
//  1. Function: takes 1 argument of type uint32 and returns (bool, error)
//  2. slice of type []uint32
//
// Returns:
//  new filtered list and error
func FilterUint32Err(f func(uint32) (bool, error), list []uint32) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	var newList []uint32
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterUint16Err takes two arguments
//  1. Function: takes 1 argument of type uint16 and returns (bool, error)
//  2. slice of type []uint16
//
// Returns:
//  new filtered list and error
func FilterUint16Err(f func(uint16) (bool, error), list []uint16) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	var newList []uint16
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterUint8Err takes two arguments
//  1. Function: takes 1 argument of type uint8 and returns (bool, error)
//  2. slice of type []uint8
//
// Returns:
//  new filtered list and error
func FilterUint8Err(f func(uint8) (bool, error), list []uint8) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	var newList []uint8
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterStrErr takes two arguments
//  1. Function: takes 1 argument of type string and returns (bool, error)
//  2. slice of type []string
//
// Returns:
//  new filtered list and error
func FilterStrErr(f func(string) (bool, error), list []string) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	var newList []string
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterBoolErr takes two arguments
//  1. Function: takes 1 argument of type bool and returns (bool, error)
//  2. slice of type []bool
//
// Returns:
//  new filtered list and error
func FilterBoolErr(f func(bool) (bool, error), list []bool) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	var newList []bool
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterFloat32Err takes two arguments
//  1. Function: takes 1 argument of type float32 and returns (bool, error)
//  2. slice of type []float32
//
// Returns:
//  new filtered list and error
func FilterFloat32Err(f func(float32) (bool, error), list []float32) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	var newList []float32
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// FilterFloat64Err takes two arguments
//  1. Function: takes 1 argument of type float64 and returns (bool, error)
//  2. slice of type []float64
//
// Returns:
//  new filtered list and error
func FilterFloat64Err(f func(float64) (bool, error), list []float64) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	var newList []float64
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if r {
			newList = append(newList, v)
		}
	}
	return newList, nil
}
