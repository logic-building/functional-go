package fp

// EveryIntPtr returns true if supplied function returns logical true for every item in the list
func EveryIntPtr(f func(*int) bool, list []*int) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryInt64Ptr returns true if supplied function returns logical true for every item in the list
func EveryInt64Ptr(f func(*int64) bool, list []*int64) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryInt32Ptr returns true if supplied function returns logical true for every item in the list
func EveryInt32Ptr(f func(*int32) bool, list []*int32) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryInt16Ptr returns true if supplied function returns logical true for every item in the list
func EveryInt16Ptr(f func(*int16) bool, list []*int16) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryInt8Ptr returns true if supplied function returns logical true for every item in the list
func EveryInt8Ptr(f func(*int8) bool, list []*int8) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUintPtr returns true if supplied function returns logical true for every item in the list
func EveryUintPtr(f func(*uint) bool, list []*uint) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUint64Ptr returns true if supplied function returns logical true for every item in the list
func EveryUint64Ptr(f func(*uint64) bool, list []*uint64) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUint32Ptr returns true if supplied function returns logical true for every item in the list
func EveryUint32Ptr(f func(*uint32) bool, list []*uint32) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUint16Ptr returns true if supplied function returns logical true for every item in the list
func EveryUint16Ptr(f func(*uint16) bool, list []*uint16) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUint8Ptr returns true if supplied function returns logical true for every item in the list
func EveryUint8Ptr(f func(*uint8) bool, list []*uint8) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryStrPtr returns true if supplied function returns logical true for every item in the list
func EveryStrPtr(f func(*string) bool, list []*string) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryBoolPtr returns true if supplied function returns logical true for every item in the list
func EveryBoolPtr(f func(*bool) bool, list []*bool) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryFloat32Ptr returns true if supplied function returns logical true for every item in the list
func EveryFloat32Ptr(f func(*float32) bool, list []*float32) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryFloat64Ptr returns true if supplied function returns logical true for every item in the list
func EveryFloat64Ptr(f func(*float64) bool, list []*float64) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
