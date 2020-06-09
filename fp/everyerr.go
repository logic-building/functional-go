package fp

// EveryIntErr returns true if supplied function returns logical true for every item in the list
func EveryIntErr(f func(int) (bool, error), list []int) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryInt64Err returns true if supplied function returns logical true for every item in the list
func EveryInt64Err(f func(int64) (bool, error), list []int64) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryInt32Err returns true if supplied function returns logical true for every item in the list
func EveryInt32Err(f func(int32) (bool, error), list []int32) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryInt16Err returns true if supplied function returns logical true for every item in the list
func EveryInt16Err(f func(int16) (bool, error), list []int16) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryInt8Err returns true if supplied function returns logical true for every item in the list
func EveryInt8Err(f func(int8) (bool, error), list []int8) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryUintErr returns true if supplied function returns logical true for every item in the list
func EveryUintErr(f func(uint) (bool, error), list []uint) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryUint64Err returns true if supplied function returns logical true for every item in the list
func EveryUint64Err(f func(uint64) (bool, error), list []uint64) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryUint32Err returns true if supplied function returns logical true for every item in the list
func EveryUint32Err(f func(uint32) (bool, error), list []uint32) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryUint16Err returns true if supplied function returns logical true for every item in the list
func EveryUint16Err(f func(uint16) (bool, error), list []uint16) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryUint8Err returns true if supplied function returns logical true for every item in the list
func EveryUint8Err(f func(uint8) (bool, error), list []uint8) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryStrErr returns true if supplied function returns logical true for every item in the list
func EveryStrErr(f func(string) (bool, error), list []string) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryBoolErr returns true if supplied function returns logical true for every item in the list
func EveryBoolErr(f func(bool) (bool, error), list []bool) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryFloat32Err returns true if supplied function returns logical true for every item in the list
func EveryFloat32Err(f func(float32) (bool, error), list []float32) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}

// EveryFloat64Err returns true if supplied function returns logical true for every item in the list
func EveryFloat64Err(f func(float64) (bool, error), list []float64) (bool, error) {
	if f == nil || len(list) == 0 {
		return false, nil
	}
	for _, v := range list {
		r, err := f(v)
		if err != nil {
			return false, err
		}
		if !r {
			return false, nil
		}
	}
	return true, nil
}
