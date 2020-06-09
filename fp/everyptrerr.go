package fp

// EveryIntPtrErr returns true if supplied function returns logical true for every item in the list
func EveryIntPtrErr(f func(*int) (bool, error), list []*int) (bool, error) {
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

// EveryInt64PtrErr returns true if supplied function returns logical true for every item in the list
func EveryInt64PtrErr(f func(*int64) (bool, error), list []*int64) (bool, error) {
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

// EveryInt32PtrErr returns true if supplied function returns logical true for every item in the list
func EveryInt32PtrErr(f func(*int32) (bool, error), list []*int32) (bool, error) {
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

// EveryInt16PtrErr returns true if supplied function returns logical true for every item in the list
func EveryInt16PtrErr(f func(*int16) (bool, error), list []*int16) (bool, error) {
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

// EveryInt8PtrErr returns true if supplied function returns logical true for every item in the list
func EveryInt8PtrErr(f func(*int8) (bool, error), list []*int8) (bool, error) {
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

// EveryUintPtrErr returns true if supplied function returns logical true for every item in the list
func EveryUintPtrErr(f func(*uint) (bool, error), list []*uint) (bool, error) {
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

// EveryUint64PtrErr returns true if supplied function returns logical true for every item in the list
func EveryUint64PtrErr(f func(*uint64) (bool, error), list []*uint64) (bool, error) {
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

// EveryUint32PtrErr returns true if supplied function returns logical true for every item in the list
func EveryUint32PtrErr(f func(*uint32) (bool, error), list []*uint32) (bool, error) {
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

// EveryUint16PtrErr returns true if supplied function returns logical true for every item in the list
func EveryUint16PtrErr(f func(*uint16) (bool, error), list []*uint16) (bool, error) {
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

// EveryUint8PtrErr returns true if supplied function returns logical true for every item in the list
func EveryUint8PtrErr(f func(*uint8) (bool, error), list []*uint8) (bool, error) {
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

// EveryStrPtrErr returns true if supplied function returns logical true for every item in the list
func EveryStrPtrErr(f func(*string) (bool, error), list []*string) (bool, error) {
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

// EveryBoolPtrErr returns true if supplied function returns logical true for every item in the list
func EveryBoolPtrErr(f func(*bool) (bool, error), list []*bool) (bool, error) {
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

// EveryFloat32PtrErr returns true if supplied function returns logical true for every item in the list
func EveryFloat32PtrErr(f func(*float32) (bool, error), list []*float32) (bool, error) {
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

// EveryFloat64PtrErr returns true if supplied function returns logical true for every item in the list
func EveryFloat64PtrErr(f func(*float64) (bool, error), list []*float64) (bool, error) {
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
