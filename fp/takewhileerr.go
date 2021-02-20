package fp

// TakeWhileIntErr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type int and returns ([]int, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileIntErr(f func(int) (bool, error), list []int) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}
	var newList []int
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileInt64Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type int64 and returns ([]int64, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileInt64Err(f func(int64) (bool, error), list []int64) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}
	var newList []int64
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileInt32Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type int32 and returns ([]int32, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileInt32Err(f func(int32) (bool, error), list []int32) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}
	var newList []int32
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileInt16Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type int16 and returns ([]int16, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileInt16Err(f func(int16) (bool, error), list []int16) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}
	var newList []int16
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileInt8Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type int8 and returns ([]int8, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileInt8Err(f func(int8) (bool, error), list []int8) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}
	var newList []int8
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileUintErr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type uint and returns ([]uint, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUintErr(f func(uint) (bool, error), list []uint) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}
	var newList []uint
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileUint64Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type uint64 and returns ([]uint64, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUint64Err(f func(uint64) (bool, error), list []uint64) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}
	var newList []uint64
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileUint32Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type uint32 and returns ([]uint32, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUint32Err(f func(uint32) (bool, error), list []uint32) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}
	var newList []uint32
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileUint16Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type uint16 and returns ([]uint16, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUint16Err(f func(uint16) (bool, error), list []uint16) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}
	var newList []uint16
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileUint8Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type uint8 and returns ([]uint8, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileUint8Err(f func(uint8) (bool, error), list []uint8) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}
	var newList []uint8
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileStrErr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type string and returns ([]string, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileStrErr(f func(string) (bool, error), list []string) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}
	var newList []string
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileBoolErr returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type bool and returns ([]bool, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileBoolErr(f func(bool) (bool, error), list []bool) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}
	var newList []bool
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileFloat32Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type float32 and returns ([]float32, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileFloat32Err(f func(float32) (bool, error), list []float32) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}
	var newList []float32
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}

// TakeWhileFloat64Err returns new list based on condition in the supplied function. It returns new list once condition fails.
//
// Takes 2 inputs:
//	1. Function - 1 input of type float64 and returns ([]float64, error)
//	2. List
//
// Returns:
//	New List, error.
//	Empty list if all the parameters are nil or either of one parameter is nil
func TakeWhileFloat64Err(f func(float64) (bool, error), list []float64) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}
	var newList []float64
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		if !r {
			return newList, nil
		}
		newList = append(newList, v)
	}
	return newList, nil
}
